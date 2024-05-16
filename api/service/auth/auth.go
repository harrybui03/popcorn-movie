package auth

import (
	"PopcornMovie/cmd/middleware"
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/gateway/email"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"time"
)

const CodeResetPassword = "code-reset-password"

type Service interface {
	Register(ctx context.Context, registerInput model.RegisterInput) (*ent.User, error)
	Login(ctx context.Context, loginInput model.LoginInput) (*model.Jwt, error)
	RenewAccessToken(ctx context.Context, input model.RenewAccessTokenInput) (*model.Jwt, error)
	ChangePassword(ctx context.Context, input model.ChangePasswordInput) (string, error)
	ForgotPassword(ctx context.Context, email string) (string, error)
	ResetPassword(ctx context.Context, input model.ResetPasswordInput) (string, error)
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.Configurations
	mailer     email.MailSender
}

func (i impl) ForgotPassword(ctx context.Context, email string) (string, error) {
	exist, _ := i.repository.User().FindUserByEmail(ctx, email)
	if exist == nil {
		return "", utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorNotFound), "User"), utils.ErrorCodeNotFound)
	}

	code := utils.GenerateNumber()
	codeReset := strconv.Itoa(int(code))

	bodyHTML := fmt.Sprintf(`
		<p>Hello %s,</p>
		<p>You have requested to reset your password.</p>
		<p>Here the code reset password ( Please don't let others know this code: %s</p>
		<br>
		<p>Ignore this email if you do remember your password, or you have not made the request.</p>
	`, email, codeReset)

	ctx = context.WithValue(ctx, CodeResetPassword, codeReset)

	err := i.mailer.SendMail(email, "Reset Password", bodyHTML)
	if err != nil {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return "Please check your email to reset password", nil
}

func (i impl) ResetPassword(ctx context.Context, input model.ResetPasswordInput) (string, error) {
	// get code from context
	codeReset, ok := ctx.Value(CodeResetPassword).(string)
	if !ok {
		return "", utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	if codeReset != input.Code {
		return "", utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorMessageNotEqual), "Code reset input", "Code reset server"), utils.ErrorCodeUnauthorized)
	}

	user, _ := i.repository.User().FindUserByEmail(ctx, input.Email)
	if user == nil {
		return "", utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorNotFound), "User"), utils.ErrorCodeNotFound)
	}

	hashPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	_, err = user.Update().SetPassword(hashPassword).Save(ctx)
	if err != nil {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return "Password reset successfully", nil
}

func (i impl) Register(ctx context.Context, registerInput model.RegisterInput) (*ent.User, error) {
	var validateInput bool
	// validate email
	validateInput = utils.ValidateEmail(registerInput.Email)
	if !validateInput {
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInvalidEmail), utils.ErrorCodeBadRequest)
	}

	// check exist
	checkExistUser, _ := i.repository.User().FindUserByEmail(ctx, registerInput.Email)
	if checkExistUser != nil {
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorInUse), "email"), utils.ErrorCodeNotFound)
	}

	// validate password
	validateInput = utils.ValidatePassword(registerInput.Password)
	if !validateInput {
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInvalidPassword), utils.ErrorCodeBadRequest)
	}

	validateInput = registerInput.Password == registerInput.ConfirmPassword
	if !validateInput {
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorMessageNotEqual), "password", "confirm password"), utils.ErrorCodeNotFound)
	}

	// hash password
	hashPassword, err := utils.HashPassword(registerInput.Password)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}
	role := model.RoleCustomer
	// Create User
	userRecord, err := i.repository.User().Create(ctx, model.CreateUserInput{
		Email:       registerInput.Email,
		Password:    hashPassword,
		DisplayName: registerInput.DisplayName,
		Role:        &role,
	})
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return userRecord, nil
}

func (i impl) Login(ctx context.Context, loginInput model.LoginInput) (*model.Jwt, error) {
	// Check exist
	user, err := i.repository.User().FindUserByEmail(ctx, loginInput.Email)

	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorIncorrect), "email"), utils.ErrorCodeNotFound)
	}
	// validate password
	err = utils.ComparePassword(user.Password, loginInput.Password)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorIncorrect), "password"), utils.ErrorCodeBadRequest)
	}
	// gen JWT
	accessToken, err, _ := utils.CreateToken(user.ID, string(user.Role), i.appConfig.AppConfig.JWTDuration, i.appConfig.AppConfig.JWTSecret)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	// gen Refresh token
	refreshToken, err, payload := utils.CreateToken(user.ID, string(user.Role), i.appConfig.AppConfig.RefreshTokenDuration, i.appConfig.AppConfig.JWTSecret)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	_, err = i.repository.Session().Create(ctx, model.CreateSessionInput{
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(i.appConfig.AppConfig.RefreshTokenDuration),
		UserID:       user.ID.String(),
		ID:           payload.ID.String(),
	})
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return &model.Jwt{
		AccessToken:  "Bearer " + accessToken,
		RefreshToken: "Bearer " + refreshToken,
	}, nil
}

func (i impl) RenewAccessToken(ctx context.Context, input model.RenewAccessTokenInput) (*model.Jwt, error) {
	payload, err := utils.VerifyToken(input.RefreshToken, i.appConfig.AppConfig.JWTSecret)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	getSession, err := i.repository.Session().GetSessionByID(ctx, payload.ID)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	if getSession.RefreshToken != input.RefreshToken {
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorMessageNotEqual), "refresh token ", "refresh token "), utils.ErrorCodeUnauthorized)
	}

	if time.Now().After(getSession.ExpiresAt) {
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorUnauthorizedRequest), utils.ErrorCodeUnauthorized)
	}

	refreshToken, err, _ := utils.CreateToken(payload.UserID, payload.Role, i.appConfig.AppConfig.RefreshTokenDuration, i.appConfig.AppConfig.JWTSecret)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return &model.Jwt{
		RefreshToken: "Bearer " + refreshToken,
	}, nil
}

func (i impl) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (string, error) {
	payload := middleware.GetPayload(ctx)
	user, err := i.repository.User().FindUserByID(ctx, payload.UserID)
	if err != nil {
		i.logger.Error(err.Error())
		return "", err
	}

	err = utils.ComparePassword(user.Password, input.OldPassword)
	if err != nil {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorMessageNotEqual), "Old password ", "Current Password"), utils.ErrorCodeUnauthorized)
	}

	if input.ConfirmNewPassword != input.NewPassword {
		return "", utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorMessageNotEqual), "New password ", "Confirm Password"), utils.ErrorCodeUnauthorized)
	}

	newPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	_, err = user.Update().SetPassword(newPassword).Save(ctx)
	if err != nil {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return "Password change successfully", nil

}

func New(repository repository.Registry, logger *zap.Logger, mailer email.MailSender, appConfig config.Configurations) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
		mailer:     mailer,
	}
}

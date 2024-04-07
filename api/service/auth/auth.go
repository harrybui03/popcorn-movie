package auth

import (
	"PopcornMovie/cmd/middleware"
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	Register(ctx context.Context, registerInput model.RegisterInput) (*ent.User, error)
	Login(ctx context.Context, loginInput model.LoginInput) (*model.Jwt, error)
	RenewAccessToken(ctx context.Context, input model.RenewAccessTokenInput) (*model.Jwt, error)
	ChangePassword(ctx context.Context, input model.ChangePasswordInput) (string, error)
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.AppConfig
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

	// Create User
	userRecord, err := i.repository.User().Create(ctx, model.CreateUserInput{
		Email:       registerInput.Email,
		Password:    hashPassword,
		DisplayName: registerInput.DisplayName,
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
	accessToken, err, _ := utils.CreateToken(user.ID, string(user.Role), i.appConfig.JWTDuration, i.appConfig.JWTSecret)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	// gen Refresh token
	refreshToken, err, payload := utils.CreateToken(user.ID, string(user.Role), i.appConfig.RefreshTokenDuration, i.appConfig.JWTSecret)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	_, err = i.repository.Session().Create(ctx, model.CreateSessionInput{
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(i.appConfig.RefreshTokenDuration),
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
	payload, err := utils.VerifyToken(input.RefreshToken, i.appConfig.JWTSecret)
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

	refreshToken, err, _ := utils.CreateToken(payload.UserID, payload.Role, i.appConfig.RefreshTokenDuration, i.appConfig.JWTSecret)
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

func New(repository repository.Registry, logger *zap.Logger, appConfig config.AppConfig) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
}

package auth

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type Service interface {
	Register(ctx context.Context, registerInput model.RegisterInput) (*ent.User, error)
	Login(ctx context.Context, loginInput model.LoginInput) (*model.Jwt, error)
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.AppConfig
}

func New(repository repository.Registry, logger *zap.Logger, appConfig config.AppConfig) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorIncorrect), "email"), utils.ErrorCodeNotFound)
		}
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}
	// validate password
	err = utils.ComparePassword(user.Password, loginInput.Password)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorIncorrect), "password"), utils.ErrorCodeBadRequest)
	}
	// gen JWT
	accessToken, err := utils.CreateToken(user.ID, string(user.Role), i.appConfig.JWTDuration, i.appConfig.JWTSecret)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return &model.Jwt{
		AccessToken: accessToken,
	}, nil
}

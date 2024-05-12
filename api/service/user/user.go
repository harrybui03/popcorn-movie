package user

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"fmt"
	"go.uber.org/zap"
)

// Service is the interface for the user service.
type Service interface {
	GetAllUsers(ctx context.Context, input model.ListUserInput) ([]*ent.User, int, error)
	CreateUser(ctx context.Context, input model.CreateUserInput) (*ent.User, error)
	UpdateUser(ctx context.Context, input model.UpdateUserInput) (*ent.User, error)
}

// impl is the implementation of the user service.
type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.AppConfig
}

func (i impl) CreateUser(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	var role model.Role
	// check exist
	checkExistUser, _ := i.repository.User().FindUserByEmail(ctx, input.Email)
	if checkExistUser != nil {
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorInUse), "email"), utils.ErrorCodeNotFound)
	}

	validateInput := utils.ValidatePassword(input.Password)
	if !validateInput {
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInvalidPassword), utils.ErrorCodeBadRequest)
	}

	// hash password
	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	if input.Role != nil {
		role = *input.Role
	}

	// Create User
	userRecord, err := i.repository.User().Create(ctx, model.CreateUserInput{
		Email:       input.Email,
		Password:    hashPassword,
		DisplayName: input.DisplayName,
		Role:        &role,
	})

	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return userRecord, nil
}

func (i impl) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*ent.User, error) {
	//// Get User by ID
	//userId, err := uuid.Parse(input.ID)
	//if err != nil {
	//	i.logger.Error(err.Error())
	//	return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	//}
	//var userRecord *ent.User
	//userRecord, _ = i.repository.User().FindUserByID(ctx, userId)
	//if userRecord == nil {
	//	i.logger.Error(err.Error())
	//	return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorNotFound), "user"), utils.ErrorCodeNotFound)
	//}
	//
	////updateQuery := userRecord.Update()
	//
	//// Update User
	//if input.Email != nil {
	//	if *input.Email != userRecord.Email {
	//		// check exist
	//		checkExistUser, _ := i.repository.User().FindUserByEmail(ctx, *input.Email)
	//		if checkExistUser != nil {
	//			return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorInUse), "email"), utils.ErrorCodeNotFound)
	//		}
	//		updateQuery = userRecord.Update().SetEmail(*input.Email)
	//	}
	//}
	//
	//if input.DisplayName != nil {
	//	if *input.DisplayName != userRecord.Displayname {
	//		updateQuery = userRecord.Update().SetDisplayname(*input.DisplayName)
	//	}
	//}
	//
	//if input.Role != nil {
	//	if user.Role(*input.Role) != userRecord.Role {
	//		updateQuery = userRecord.Update().SetRole(user.Role(*input.Role))
	//	}
	//}
	//
	//if input.IsLocked != nil {
	//	if *input.IsLocked != userRecord.IsLocked {
	//		updateQuery = userRecord.Update().SetIsLocked(*input.IsLocked)
	//	}
	//}
	//
	//userRecordUpdated, err := userRecord.Update().Save(ctx)
	//if err != nil {
	//	return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	//}

	return nil, nil
}

func (i impl) GetAllUsers(ctx context.Context, input model.ListUserInput) ([]*ent.User, int, error) {
	query, err := i.repository.User().UserQuery(ctx)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	count, err := i.repository.User().CountUsers(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	if input.Pagination != nil {
		offset := utils.CalculateOffset(input.Pagination.Page, input.Pagination.Limit)
		query.Limit(input.Pagination.Limit).Offset(offset)
	}

	users, err := i.repository.User().GetAllUsers(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return users, count, nil
}

// New creates a new user service.
func New(repository repository.Registry, logger *zap.Logger, appConfig config.AppConfig) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
}

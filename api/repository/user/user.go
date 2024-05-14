package user

import (
	"PopcornMovie/ent"
	"PopcornMovie/ent/user"
	"PopcornMovie/model"
	"context"
	"github.com/google/uuid"
)

// Repository is the interface for the user repository.
type Repository interface {
	// Create creates a new user.
	Create(ctx context.Context, input model.CreateUserInput) (*ent.User, error)

	UpdateQuery() *ent.UserUpdate

	FindUserByEmail(ctx context.Context, email string) (*ent.User, error)

	FindUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error)

	UserQuery(ctx context.Context) (*ent.UserQuery, error)

	GetAllUsers(ctx context.Context, query *ent.UserQuery) ([]*ent.User, error)

	CountUsers(ctx context.Context, query *ent.UserQuery) (int, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) UpdateQuery() *ent.UserUpdate {
	return i.client.User.Update()
}

func (i impl) CountUsers(ctx context.Context, query *ent.UserQuery) (int, error) {
	return query.Count(ctx)
}

func (i impl) UserQuery(ctx context.Context) (*ent.UserQuery, error) {
	return i.client.User.Query(), nil
}

func (i impl) GetAllUsers(ctx context.Context, query *ent.UserQuery) ([]*ent.User, error) {
	return query.All(ctx)
}

func (i impl) Create(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	userRecord, err := i.client.User.
		Create().
		SetDisplayname(input.DisplayName).
		SetEmail(input.Email).
		SetPassword(input.Password).
		SetRole(user.Role(*input.Role)).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return userRecord, nil
}

func (i impl) FindUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	userRecord, err := i.client.User.Query().Where(user.Email(email)).First(ctx)
	if err != nil {
		return nil, err
	}

	return userRecord, err
}

func (i impl) FindUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	userRecord, err := i.client.User.Query().Where(user.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	return userRecord, err
}

// New creates a new user repository.
func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}

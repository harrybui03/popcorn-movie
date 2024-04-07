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

	FindUserByEmail(ctx context.Context, email string) (*ent.User, error)

	FindUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) Create(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	userRecord, err := i.client.User.
		Create().
		SetDisplayname(input.DisplayName).
		SetEmail(input.Email).
		SetPassword(input.Password).
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

package session

import (
	"PopcornMovie/ent"
	session2 "PopcornMovie/ent/session"
	"PopcornMovie/model"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Repository interface {
	Create(ctx context.Context, input model.CreateSessionInput) (*ent.Session, error)
	GetSessionByID(ctx context.Context, id uuid.UUID) (*ent.Session, error)
}

type impl struct {
	client *ent.Client
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}

func (i impl) Create(ctx context.Context, input model.CreateSessionInput) (*ent.Session, error) {
	id, _ := uuid.Parse(input.ID)
	userId, _ := uuid.Parse(input.UserID)

	session, err := i.client.Session.
		Create().
		SetID(id).
		SetUserID(userId).
		SetRefreshToken(input.RefreshToken).
		SetExpiresAt(input.ExpiresAt).
		Save(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return session, nil
}

func (i impl) GetSessionByID(ctx context.Context, id uuid.UUID) (*ent.Session, error) {
	session, err := i.client.Session.Query().Where(session2.ID(id)).First(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return session, nil
}

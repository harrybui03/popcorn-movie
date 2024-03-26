package repository

import (
	"PopcornMovie/ent"
	"PopcornMovie/repository/user"
)

// Registry is the interface for the repository registry.
type Registry interface {
	User() user.Repository
}

type impl struct {
	user user.Repository
}

// New creates a new repository registry.
func New(client *ent.Client) Registry {
	return &impl{
		user: user.New(client),
	}
}

func (i impl) User() user.Repository {
	return i.user
}

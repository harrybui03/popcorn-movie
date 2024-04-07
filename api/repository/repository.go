package repository

import (
	"PopcornMovie/ent"
	"PopcornMovie/repository/room"
	"PopcornMovie/repository/session"
	"PopcornMovie/repository/theater"
	"PopcornMovie/repository/user"
)

// Registry is the interface for the repository registry.
type Registry interface {
	User() user.Repository
	Theater() theater.Repository
	Session() session.Repository
	Room() room.Repository
}

type impl struct {
	user    user.Repository
	theater theater.Repository
	session session.Repository
	room    room.Repository
}

func (i impl) Room() room.Repository {
	return i.room
}

func (i impl) Session() session.Repository {
	return i.session
}

func (i impl) Theater() theater.Repository { return i.theater }

func (i impl) User() user.Repository {
	return i.user
}

// New creates a new repository registry.
func New(client *ent.Client) Registry {
	return &impl{
		user:    user.New(client),
		theater: theater.New(client),
		session: session.New(client),
	}
}

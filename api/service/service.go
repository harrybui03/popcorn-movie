package service

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/gateway/cloudinary"
	"PopcornMovie/gateway/email"
	"PopcornMovie/repository"
	"PopcornMovie/service/auth"
	"PopcornMovie/service/food"
	"PopcornMovie/service/movie"
	"PopcornMovie/service/room"
	"PopcornMovie/service/seat"
	"PopcornMovie/service/show_time"
	"PopcornMovie/service/theater"
	"PopcornMovie/service/ticket"
	"PopcornMovie/service/transaction"
	"PopcornMovie/service/user"
	"go.uber.org/zap"
)

type Registry interface {
	User() user.Service
	Auth() auth.Service
	Theater() theater.Service
	Room() room.Service
	Food() food.Service
	Movie() movie.Service
	ShowTime() show_time.Service
	Seat() seat.Service
	Transaction() transaction.Service
	Ticket() ticket.Service
}

type impl struct {
	user        user.Service
	auth        auth.Service
	theater     theater.Service
	room        room.Service
	food        food.Service
	movie       movie.Service
	showTime    show_time.Service
	seat        seat.Service
	transaction transaction.Service
	ticket      ticket.Service
}

func (i impl) Ticket() ticket.Service {
	return i.ticket
}

func (i impl) Transaction() transaction.Service {
	return i.transaction
}

func (i impl) Seat() seat.Service {
	return i.seat
}

func (i impl) ShowTime() show_time.Service {
	return i.showTime
}

func (i impl) Movie() movie.Service {
	return i.movie
}

func (i impl) Food() food.Service {
	return i.food
}

func (i impl) Room() room.Service {
	return i.room
}

func (i impl) Theater() theater.Service { return i.theater }

func (i impl) User() user.Service {
	//TODO implement me
	return i.user
}

func (i impl) Auth() auth.Service {
	return i.auth
}

func New(entClient *ent.Client, logger *zap.Logger, appConfig config.Configurations) Registry {
	repositoryRegistry := repository.New(entClient)
	mailer := email.New(appConfig)
	clouldinary, _ := cloudinary.New(appConfig, logger)

	return &impl{
		user:        user.New(repositoryRegistry, logger, appConfig),
		auth:        auth.New(repositoryRegistry, logger, mailer, appConfig),
		theater:     theater.New(repositoryRegistry, logger, appConfig),
		room:        room.New(repositoryRegistry, logger, appConfig),
		food:        food.New(repositoryRegistry, logger, appConfig),
		movie:       movie.New(repositoryRegistry, logger, appConfig, clouldinary),
		showTime:    show_time.New(repositoryRegistry, logger, appConfig),
		seat:        seat.New(repositoryRegistry, logger, appConfig),
		transaction: transaction.New(repositoryRegistry, logger, appConfig),
		ticket:      ticket.New(repositoryRegistry, logger, appConfig),
	}
}

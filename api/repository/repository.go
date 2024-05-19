package repository

import (
	"PopcornMovie/ent"
	"PopcornMovie/repository/food"
	"PopcornMovie/repository/food_order_line"
	"PopcornMovie/repository/movie"
	"PopcornMovie/repository/reset_password"
	"PopcornMovie/repository/room"
	"PopcornMovie/repository/seat"
	"PopcornMovie/repository/session"
	"PopcornMovie/repository/show_time"
	"PopcornMovie/repository/theater"
	"PopcornMovie/repository/ticket"
	"PopcornMovie/repository/transaction"
	"PopcornMovie/repository/user"
	"context"
	"fmt"
	"github.com/pkg/errors"
)

// Registry is the interface for the repository registry.
type Registry interface {
	User() user.Repository
	Theater() theater.Repository
	Session() session.Repository
	Room() room.Repository
	Food() food.Repository
	Movie() movie.Repository
	ShowTime() show_time.Repository
	Seat() seat.Repository
	Transaction() transaction.Repository
	Ticket() ticket.Repository
	FoodOrderLine() food_order_line.Repository
	ResetPassword() reset_password.Repository
	DoinTx(ctx context.Context, txFunc func(ctx context.Context, repoRegistry Registry) error) error
}

type impl struct {
	user          user.Repository
	theater       theater.Repository
	session       session.Repository
	room          room.Repository
	food          food.Repository
	movie         movie.Repository
	showTime      show_time.Repository
	seat          seat.Repository
	ticket        ticket.Repository
	foodOrderLine food_order_line.Repository
	transaction   transaction.Repository
	resetPassword reset_password.Repository
	client        *ent.Client
	tx            *ent.Tx
}

func (i impl) DoinTx(ctx context.Context, txFunc func(ctx context.Context, repoRegistry Registry) error) error {
	tx, err := i.client.Tx(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	commited := false
	defer func() {
		if commited {
			return
		}

		_ = tx.Rollback()
	}()

	impl := &impl{
		tx:            tx,
		user:          user.New(tx.Client()),
		theater:       theater.New(tx.Client()),
		session:       session.New(tx.Client()),
		room:          room.New(tx.Client()),
		food:          food.New(tx.Client()),
		movie:         movie.New(tx.Client()),
		showTime:      show_time.New(tx.Client()),
		seat:          seat.New(tx.Client()),
		ticket:        ticket.New(tx.Client()),
		foodOrderLine: food_order_line.New(tx.Client()),
		transaction:   transaction.New(tx.Client()),
		resetPassword: reset_password.New(tx.Client()),
	}

	if err := txFunc(ctx, impl); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(fmt.Errorf("failed to commit tx: %s", err.Error()))
	}

	commited = true

	return nil
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

func (i impl) Food() food.Repository {
	return i.food
}

func (i impl) Movie() movie.Repository {
	return i.movie
}

func (i impl) ShowTime() show_time.Repository {
	return i.showTime
}

func (i impl) Seat() seat.Repository {
	return i.seat
}

func (i impl) Transaction() transaction.Repository {
	return i.transaction
}

func (i impl) Ticket() ticket.Repository {
	return i.ticket
}

func (i impl) FoodOrderLine() food_order_line.Repository {
	return i.foodOrderLine
}

func (i impl) ResetPassword() reset_password.Repository {
	return i.resetPassword
}

// New creates a new repository registry.
func New(client *ent.Client) Registry {
	return &impl{
		user:          user.New(client),
		theater:       theater.New(client),
		session:       session.New(client),
		room:          room.New(client),
		food:          food.New(client),
		movie:         movie.New(client),
		showTime:      show_time.New(client),
		seat:          seat.New(client),
		ticket:        ticket.New(client),
		foodOrderLine: food_order_line.New(client),
		transaction:   transaction.New(client),
		resetPassword: reset_password.New(client),
		client:        client,
	}
}

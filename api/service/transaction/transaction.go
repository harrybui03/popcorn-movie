package transaction

import (
	"PopcornMovie/cmd/middleware"
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	CreateTransaction(ctx context.Context, input model.CreateTransactionInput) (*ent.Transaction, error)
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.AppConfig
}

func (i impl) CreateTransaction(ctx context.Context, input model.CreateTransactionInput) (*ent.Transaction, error) {
	var (
		transaction *ent.Transaction
		err         error
	)

	err = i.repository.DoinTx(ctx, func(ctx context.Context, repo repository.Registry) error {
		userId := middleware.GetPayload(ctx).UserID
		// Create Transaction
		transaction, err = repo.Transaction().CreateTransaction(ctx, userId)
		if err != nil {
			return err
		}

		showTimeID, err := uuid.Parse(input.ShowTimeID)
		if err != nil {
			return err
		}

		total := 0.0
		// Create Ticket
		for _, ticket := range input.SeatIDs {
			ticketID, err := uuid.Parse(ticket.SeatID)
			if err != nil {
				return err
			}

			_, err = repo.Ticket().CreateTicket(ctx, model.CreateTicket{
				SeatID:        ticketID,
				TransactionID: transaction.ID,
				ShowTimeID:    showTimeID,
				Price:         ticket.Price,
				IsBooked:      true,
			})

			total += ticket.Price

			if err != nil {
				return err
			}
		}

		// Create Food Order
		for _, foodOrder := range input.Foods {
			foodID, err := uuid.Parse(foodOrder.FoodID)
			if err != nil {
				return err
			}

			food, err := repo.Food().GetFoodByID(ctx, foodID)
			if err != nil {
				return err
			}

			foodOrderRecord, err := repo.FoodOrderLine().CreateFoodOrder(ctx, model.CreateFoodOrderLine{
				FoodID:        foodID,
				Quantity:      foodOrder.Quantity,
				TransactionID: transaction.ID,
			})

			total += float64(foodOrderRecord.Quantity) * food.Price
		}
		// Call VNPAY API

		// Send Email to User

		// Update Transaction
		transaction.Total = total
		transaction.Update().SaveX(ctx)

		return nil
	})

	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return transaction, err
}

func New(repository repository.Registry, logger *zap.Logger, appConfig config.AppConfig) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
}

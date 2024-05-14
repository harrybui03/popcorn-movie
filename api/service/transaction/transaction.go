package transaction

import (
	"PopcornMovie/cmd/middleware"
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/ent/ticket"
	"PopcornMovie/ent/transaction"
	"PopcornMovie/ent/user"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"github.com/google/uuid"
	"github.com/payOSHQ/payos-lib-golang"
	"go.uber.org/zap"
)

type Service interface {
	CreateTransaction(ctx context.Context, input model.CreateTransactionInput) (*payos.CheckoutResponseDataType, error)
	GetAllTransactions(ctx context.Context, input model.ListTransactionInput) ([]*ent.Transaction, int, error)
	VerifyPaymentData(ctx context.Context, webhookDataReq payos.WebhookType) error
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.Configurations
}

func (i impl) VerifyPaymentData(ctx context.Context, webhookDataReq payos.WebhookType) error {
	err := i.repository.DoinTx(ctx, func(ctx context.Context, repo repository.Registry) error {
		webhookData, err := payos.VerifyPaymentWebhookData(webhookDataReq)
		orderCode := webhookData.OrderCode
		transactionRecord, err := repo.Transaction().TransactionQuery().Where(transaction.Code(int(orderCode))).First(ctx)
		if err != nil {
			return err
		}

		// SUCCESS
		if webhookData.Code == "00" {
			ticketsArray := transactionRecord.QueryTickets().AllX(ctx)
			for _, ticketRecord := range ticketsArray {
				ticketRecord.Update().SetIsBooked(true).SaveX(ctx)
			}
			transactionRecord.Update().SetStatus(transaction.StatusPAID).SaveX(ctx)
		} else {
			transactionRecord.Update().SetStatus(transaction.StatusCANCEL).SaveX(ctx)
		}

		return nil
	})

	if err != nil {
		i.logger.Error(err.Error())
		return err
	}

	return nil
}

func (i impl) GetAllTransactions(ctx context.Context, input model.ListTransactionInput) ([]*ent.Transaction, int, error) {
	query := i.repository.Transaction().TransactionQuery().WithUser()
	if input.Filter != nil {
		userId, err := uuid.Parse(input.Filter.UserID)
		if err != nil {
			i.logger.Error(err.Error())
			return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeBadRequest)
		}
		query.Where(transaction.HasUserWith(user.ID(userId)))
	}

	if input.Pagination != nil {
		offset := utils.CalculateOffset(input.Pagination.Page, input.Pagination.Limit)
		query.Limit(input.Pagination.Limit).Offset(offset)
	}

	count, err := i.repository.Transaction().CountTransactions(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	transactions, err := i.repository.Transaction().GetAllTransactions(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return transactions, count, nil
}

func (i impl) CreateTransaction(ctx context.Context, input model.CreateTransactionInput) (*payos.CheckoutResponseDataType, error) {
	var (
		output *payos.CheckoutResponseDataType
		err    error
	)

	err = i.repository.DoinTx(ctx, func(ctx context.Context, repo repository.Registry) error {
		userId := middleware.GetPayload(ctx).UserID
		items := []payos.Item{}
		total := 0.0
		// Create Transaction
		transaction, err := repo.Transaction().CreateTransaction(ctx, userId)
		if err != nil {
			return err
		}

		showTimeID, err := uuid.Parse(input.ShowTimeID)
		if err != nil {
			return err
		}

		// Update Ticket
		for _, seat := range input.TicketIDs {
			seatID, err := uuid.Parse(seat.SeatID)
			if err != nil {
				return err
			}

			// Get Ticket By ShowTimeID and SeatID
			ticketRecord, err := repo.Ticket().TicketQuery().Where(ticket.And(ticket.ShowTimeID(showTimeID), ticket.SeatID(seatID))).First(ctx)
			if err != nil {
				return err
			}

			// update ticket
			repo.Ticket().TicketUpdate().SetTransactionID(transaction.ID).Where(ticket.ID(ticketRecord.ID)).SaveX(ctx)
			total += ticketRecord.Price

			if err != nil {
				return err
			}
		}

		items = append(items, payos.Item{Name: "Vé xem phim", Price: int(total), Quantity: len(input.TicketIDs)})

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

			if err != nil {
				return err
			}

			items = append(items, payos.Item{Name: food.Name, Price: int(float64(foodOrderRecord.Quantity) * food.Price), Quantity: foodOrderRecord.Quantity})

			total += float64(foodOrderRecord.Quantity) * food.Price
		}
		orderCode := utils.GenerateNumber()
		// Call PayOS API
		body := payos.CheckoutRequestType{
			OrderCode:   orderCode,
			Amount:      int(total),
			Items:       items,
			Description: "Thanh toán đơn hàng",
			CancelUrl:   i.appConfig.Payos.Domain,
			ReturnUrl:   i.appConfig.Payos.Domain + "/success/",
		}

		_, err = transaction.Update().SetTotal(total).SetCode(int(orderCode)).Save(ctx)
		if err != nil {
			return err
		}

		data, err := payos.CreatePaymentLink(body)
		if err != nil {
			return err
		}
		output = data

		return nil
	})

	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return output, nil
}

func New(repository repository.Registry, logger *zap.Logger, appConfig config.Configurations) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
}

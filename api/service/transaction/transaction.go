package transaction

import (
	"PopcornMovie/cmd/middleware"
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/ent/ticket"
	"PopcornMovie/ent/transaction"
	"PopcornMovie/ent/user"
	"PopcornMovie/gateway/email"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/payOSHQ/payos-lib-golang"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	CreateTransaction(ctx context.Context, input model.CreateTransactionInput) (*payos.CheckoutResponseDataType, error)
	GetAllTransactions(ctx context.Context, input model.ListTransactionInput) ([]*ent.Transaction, int, error)
	VerifyPaymentData(ctx context.Context, webhookDataReq payos.WebhookType) error
	GetRevenue(ctx context.Context, input model.RevenueInput) (*model.YearlyRevenueOutput, error)
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.Configurations
	mailer     email.MailSender
}

func (i impl) GetRevenue(ctx context.Context, input model.RevenueInput) (*model.YearlyRevenueOutput, error) {
	total := 0.0
	monthlyArr := make([]*model.MonthlyRevenueOutput, 12)
	for idx := 1; idx <= 12; idx++ {
		totalInMonth, err := i.repository.Transaction().GetRevenue(ctx, input.Year, idx)
		if err != nil {
			i.logger.Error(err.Error())
			return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
		}
		monthlyArr[idx-1] = &model.MonthlyRevenueOutput{
			Month: idx,
			Total: totalInMonth,
		}

		total += totalInMonth
	}

	return &model.YearlyRevenueOutput{
		Total: total,
		Arr:   monthlyArr,
	}, nil
}

func (i impl) VerifyPaymentData(ctx context.Context, webhookDataReq payos.WebhookType) error {
	err := i.repository.DoinTx(ctx, func(ctx context.Context, repo repository.Registry) error {
		webhookData, err := payos.VerifyPaymentWebhookData(webhookDataReq)
		orderCode := webhookData.OrderCode
		transactionRecord, err := repo.Transaction().TransactionQuery().Where(transaction.Code(int(orderCode))).WithUser().WithTickets().First(ctx)
		if err != nil {
			return err
		}

		// SUCCESS
		if webhookData.Code == "00" {
			ticketsArray := transactionRecord.QueryTickets().AllX(ctx)
			for _, ticketRecord := range ticketsArray {
				_, err = ticketRecord.Update().SetIsBooked(true).SetTransactionID(transactionRecord.ID).Save(ctx)
				if err != nil {
					return err
				}
			}
			emailContent := fmt.Sprintf(`<!DOCTYPE html>
<html lang="vi">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cảm ơn bạn đã mua lựa chọn chúng tôi</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            color: #333;
            margin: 0;
            padding: 0;
        }
        .container {
            width: 100%%;
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            padding: 20px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        .header {
            text-align: center;
            padding-bottom: 20px;
        }
        .header h1 {
            margin: 0;
            color: #ff6600;
        }
        .content {
            line-height: 1.6;
        }
        .order-details {
            background-color: #f9f9f9;
            padding: 15px;
            margin: 20px 0;
            border-radius: 5px;
        }
        .order-details h3 {
            margin-top: 0;
        }
        .footer {
            text-align: center;
            padding: 10px;
            font-size: 12px;
            color: #777;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Cảm ơn bạn đã mua lựa chọn chúng tôi!</h1>
        </div>
        <div class="content">
            <p>Chào %s,</p>
            <p>Cảm ơn bạn đã mua hàng tại cửa hàng của chúng tôi. Đơn hàng của bạn đã được xử lý thành công.</p>
            <p>Lưu ý Tuyệt đối chỉ đưa cho nhân viên tại rạp mã đơn hàng để được lấy vé</p>
            <div class="order-details">
                <h3>Thông tin đơn hàng</h3>
                <p><strong>Mã đơn hàng:</strong> %d</p>
                <p><strong>Giá tiền:</strong> %d VND</p>
            </div>
            <p>Nếu bạn có bất kỳ câu hỏi nào về đơn hàng, vui lòng liên hệ với chúng tôi qua email hoặc số điện thoại hỗ trợ.</p>
            <p>Trân trọng,<br>Cửa hàng %s</p>
        </div>
        <div class="footer">
            <p>&copy; 2024 %s. All rights reserved.</p>
        </div>
    </div>
</body>
</html>`, transactionRecord.Edges.User.Displayname, orderCode, webhookData.Amount, "Popcorn Movie", "Popcorn Movie")

			err := i.mailer.SendMail(transactionRecord.Edges.User.Email, "Cảm ơn quý khách đã đặt vé tại Popcorn Movie", emailContent)
			if err != nil {
				return err
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

		if err != nil {
			return err
		}

		// Update Ticket
		for _, ticketRecord := range input.TicketIDs {
			ticketID, err := uuid.Parse(ticketRecord.ID)
			if err != nil {
				return err
			}

			// Get Ticket By ShowTimeID and SeatID
			ticketRecord, err := repo.Ticket().TicketQuery().Where(ticket.ID(ticketID)).First(ctx)
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
			CancelUrl:   i.appConfig.Payos.Domain + "/payments/cancel",
			ReturnUrl:   i.appConfig.Payos.Domain + "/payments/success/",
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

// Function to get all days in a month for a given month and year
func getAllDaysInMonth(year int, month time.Month) []time.Time {
	// Get the first day of the month
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	// Calculate the last day of the month
	lastDay := firstDay.AddDate(0, 1, -1)

	// Prepare a slice to hold all days of the month
	var days []time.Time

	// Iterate through all days in the month
	for day := firstDay; !day.After(lastDay); day = day.AddDate(0, 0, 1) {
		days = append(days, day)
	}

	return days
}

func New(repository repository.Registry, logger *zap.Logger, appConfig config.Configurations, mailer email.MailSender) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
		mailer:     mailer,
	}
}

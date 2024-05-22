package rest

import (
	"PopcornMovie/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/payOSHQ/payos-lib-golang"
)

type RestAPI interface {
	WebHooktype(c *gin.Context)
}

type impl struct {
	service service.Registry
	ctx     context.Context
}

func (i impl) WebHooktype(c *gin.Context) {
	var webhookDataReq payos.WebhookType

	if err := c.BindJSON(&webhookDataReq); err != nil {
		return
	}

	err := i.service.Transaction().VerifyPaymentData(c.Copy(), webhookDataReq)
	if err != nil {
		return
	}
}

func New(service service.Registry) RestAPI {
	return &impl{
		service: service,
	}
}

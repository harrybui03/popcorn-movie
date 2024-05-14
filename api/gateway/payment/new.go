package payment

import (
	"PopcornMovie/config"
	"github.com/payOSHQ/payos-lib-golang"
)

func NewPaymentService(config config.PayosConfig) error {
	err := payos.Key(config.ClientID, config.APIKey, config.ChecksumKey)
	if err != nil {
		return nil
	}

	return nil
}

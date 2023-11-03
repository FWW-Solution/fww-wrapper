package midtrans_payment

import (
	"fww-wrapper/internal/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func InitCore(cfg *config.MidtransConfig) *coreapi.Client {
	midtrans.ServerKey = cfg.ServerKey
	if cfg.IsProduction {
		midtrans.Environment = midtrans.Production
	} else {
		midtrans.Environment = midtrans.Sandbox
	}

	c := coreapi.Client{}

	c.New(cfg.ServerKey, midtrans.Sandbox)

	return &c

}

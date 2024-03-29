package payment

import (
	"github.com/muhammadzhuhry/bwastartup/infra"
	"github.com/muhammadzhuhry/bwastartup/user"
	"github.com/veritrans/go-midtrans"
	"log"
	"strconv"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	config, err := infra.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	midclient := midtrans.NewClient()
	midclient.ServerKey = config.ServerKey
	midclient.ClientKey = config.ClientKey
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}

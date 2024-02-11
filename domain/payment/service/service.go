package payment

import (
	"strconv"

	"github.com/agusheryanto182/go-online-store-mvp/config"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
	cnf config.Config
}

type Service interface {
	GetPaymentURL(transaction entities.Transaction, user entities.User) (string, error)
}

func NewPaymentService(cnf config.Config) *service {
	return &service{cnf: cnf}
}

func (s *service) GetPaymentURL(transaction entities.Transaction, user entities.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = s.cnf.ServerKeyMidtrans
	midclient.ClientKey = s.cnf.ClientKeyMidtrans
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Fullname,
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

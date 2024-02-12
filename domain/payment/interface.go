package payment

import (
	"github.com/agusheryanto182/go-online-store-mvp/entities"
)

type PaymentServiceInterface interface {
	GetPaymentURL(transaction entities.Transaction, user entities.User) (string, error)
}

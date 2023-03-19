package interfaces

import (
	"github.com/akshayur04/project-ecommerce/pkg/domain"
)

type PaymentRepository interface {
	ViewPaymentDetails(orderID int) (domain.PaymentDetails, error)
	UpdatePaymentDetails(orderID int, paymentRef string) (domain.PaymentDetails, error)
}

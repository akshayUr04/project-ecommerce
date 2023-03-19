package interfaces

import (
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
)

type PaymentUseCase interface {
	CreateRazorpayPayment(userId, orderId int) (domain.Orders, string, error)
	UpdatePaymentDetails(paymentVerifier helperStruct.PaymentVerification) error
}

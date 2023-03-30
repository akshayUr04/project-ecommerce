package usecase

import (
	"fmt"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/config"
	"github.com/akshayur04/project-ecommerce/pkg/domain"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/razorpay/razorpay-go"
)

type PaymentUseCase struct {
	paymentRepo interfaces.PaymentRepository
	orderRepo   interfaces.OrderRepository
	cfg         config.Config
}

func NewPaymentuseCase(paymentRepo interfaces.PaymentRepository, orderRepo interfaces.OrderRepository, cfg config.Config) services.PaymentUseCase {
	return &PaymentUseCase{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
		cfg:         cfg,
	}
}

func (c *PaymentUseCase) CreateRazorpayPayment(userId, orderId int) (domain.Orders, string, error) {
	paymentDetails, err := c.paymentRepo.ViewPaymentDetails(orderId)
	if err != nil {
		return domain.Orders{}, "", err
	}
	if paymentDetails.PaymentStatusID == 3 {
		return domain.Orders{}, "", fmt.Errorf("payment already completed")
	}
	//fetch order details from the db
	order, err := c.orderRepo.ListOrder(userId, orderId)
	if err != nil {
		return domain.Orders{}, "", err
	}
	if order.Id == 0 {
		return domain.Orders{}, "", fmt.Errorf("no such order found")
	}
	client := razorpay.NewClient(c.cfg.RAZORPAYID, c.cfg.RAZORPAYSECRET)

	data := map[string]interface{}{
		"amount":   order.OrderTotal * 100,
		"currency": "INR",
		"receipt":  "test_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return domain.Orders{}, "", err
	}

	value := body["id"]
	razorpayID := value.(string)
	return order, razorpayID, err
}

func (c *PaymentUseCase) UpdatePaymentDetails(paymentVerifier helperStruct.PaymentVerification) error {
	paymentDetails, err := c.paymentRepo.ViewPaymentDetails(paymentVerifier.OrderID)
	if err != nil {
		return err
	}
	if paymentDetails.ID == 0 {
		return fmt.Errorf("no order found")
	}

	if paymentDetails.OrderTotal != paymentVerifier.Total {
		return fmt.Errorf("payment amount and order amount does not match")
	}
	updatedPayment, err := c.paymentRepo.UpdatePaymentDetails(paymentVerifier.OrderID, paymentVerifier.PaymentRef)
	if err != nil {
		return err
	}
	if updatedPayment.ID == 0 {
		return fmt.Errorf("failed to update payment details")
	}
	return nil
}

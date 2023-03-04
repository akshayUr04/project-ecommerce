package interfaces

import (
	"context"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type OtpUseCase interface {
	SendOtp(ctx context.Context, phno helperStruct.OTPData) error
	ValidateOtp(otpDetails helperStruct.VerifyOtp) (*openapi.VerifyV2VerificationCheck, error)
}

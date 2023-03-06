package usecase

import (
	"context"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/config"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type OtpUseCase struct {
	// otpRepo interfaces.OtpRepository
	cfg config.Config
}

func NewOtpUseCase(cfg config.Config) services.OtpUseCase {
	return &OtpUseCase{
		// otpRepo: repo,
		cfg: cfg,
	}
}

func (c *OtpUseCase) SendOtp(ctx context.Context, phno helperStruct.OTPData) error {
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.cfg.TWILIOACCOUNTSID,
		Password: c.cfg.TWILIOAUTHTOKEN,
	})

	params := &openapi.CreateVerificationParams{}
	params.SetTo(phno.PhoneNumber)
	params.SetChannel("sms")
	_, err := client.VerifyV2.CreateVerification(c.cfg.TWILIOSERVICESID, params)
	return err
}

func (c *OtpUseCase) ValidateOtp(otpDetails helperStruct.VerifyOtp) (*openapi.VerifyV2VerificationCheck, error) {
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.cfg.TWILIOACCOUNTSID,
		Password: c.cfg.TWILIOAUTHTOKEN,
	})
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(otpDetails.User.PhoneNumber)
	params.SetCode(otpDetails.Code)
	resp, err := client.VerifyV2.CreateVerificationCheck(c.cfg.TWILIOSERVICESID, params)
	return resp, err
}

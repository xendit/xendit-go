package main

import (
	"log"

	"github.com/xendit/xendit-go/pmsv2"
	"github.com/xendit/xendit-go/pmsv2/constant"
	"github.com/xendit/xendit-go/pmsv2/ewallet"
)

func paymentmethodTest() {
	createData := pmsv2.CreatePaymentMethodParams{
		Type:        constant.PaymentMethodTypeEwallet,
		Reusability: constant.ReusabilityOneTimeUse,
		Ewallet: &ewallet.CreateMethod{
			ChannelCode: ewallet.DANA,
			ChannelProperties: ewallet.ChannelProperties{
				"success_return_url": "https://redirect.me/goodstuff",
				"failure_return_url": "https://redirect.me/badstuff",
				"cancel_return_url":  "https://redirect.me/nostuff",
			},
		},
	}

	_, err := pmsv2.CreatePaymentMethod(&createData)
	if err != nil {
		log.Panic(err)
	}

}

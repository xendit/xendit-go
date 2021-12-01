package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

func ewalletTest() {
	createPaymentData := ewallet.CreatePaymentParams{
		ExternalID:  "ovo-" + time.Now().String(),
		Amount:      1,
		Phone:       "08123123123",
		EWalletType: xendit.EWalletTypeOVO,
	}

	resp, err := ewallet.CreatePayment(&createPaymentData)
	if err != nil {
		log.Panic(err)
	}
	getPaymentStatusData := ewallet.GetPaymentStatusParams{
		ExternalID:  resp.ExternalID,
		EWalletType: resp.EWalletType,
	}

	_, err = ewallet.GetPaymentStatus(&getPaymentStatusData)
	if err != nil {
		log.Panic(err)
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	ewalletBasketItem := xendit.EWalletBasketItem{
		ReferenceID: "basket-product-ref-id",
		Name:        "product name",
		Category:    "mechanics",
		Currency:    "IDR",
		Price:       50000,
		Quantity:    5,
		Type:        "type",
		SubCategory: "subcategory",
		Metadata:    metadata,
	}

	createEWalletChargeData := ewallet.CreateEWalletChargeParams{
		ReferenceID:    "test-reference-id",
		Currency:       "IDR",
		Amount:         1688,
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_SHOPEEPAY",
		ChannelProperties: map[string]string{
			"success_redirect_url": "https://yourwebsite.com/order/123",
		},
		Basket: []xendit.EWalletBasketItem{
			ewalletBasketItem,
		},
		Metadata: metadata,
	}

	charge, chargeErr := ewallet.CreateEWalletCharge(&createEWalletChargeData)
	if chargeErr != nil {
		log.Panic(chargeErr)
	}

	getEWalletChargeStatusData := ewallet.GetEWalletChargeStatusParams{
		ChargeID: charge.ID,
	}

	_, chargeErr = ewallet.GetEWalletChargeStatus(&getEWalletChargeStatusData)
	if chargeErr != nil {
		log.Panic(chargeErr)
	}

	fmt.Println("EWallet integration tests done!")
}

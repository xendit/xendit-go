package ewallet_test

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

func ExampleCreatePayment() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := ewallet.CreatePaymentParams{
		ExternalID:  "dana-ewallet",
		Amount:      20000,
		Phone:       "08123123123",
		EWalletType: xendit.EWalletTypeDANA,
		CallbackURL: "mystore.com/callback",
		RedirectURL: "mystore.com/redirect",
	}

	resp, err := ewallet.CreatePayment(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment: %+v\n", resp)
}

func ExampleGetPaymentStatus() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := ewallet.GetPaymentStatusParams{
		ExternalID:  "data-ewallet",
		EWalletType: xendit.EWalletTypeDANA,
	}

	resp, err := ewallet.GetPaymentStatus(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved payment: %+v\n", resp)
}

func ExampleCreateEWalletCharge() {
	xendit.Opt.SecretKey = "examplesecretkey"

	ewalletBasketItem := xendit.EWalletBasketItem{
		ReferenceID: "basket-product-ref-id",
		Name:        "product name",
		Category:    "mechanics",
		Currency:    "IDR",
		Price:       50000,
		Quantity:    5,
		Type:        "type",
		SubCategory: "subcategory",
		Metadata: map[string]interface{}{
			"meta": "data",
		},
	}

	data := ewallet.CreateEWalletChargeParams{
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
		Metadata: map[string]interface{}{
			"meta": "data",
		},
	}

	charge, chargeErr := ewallet.CreateEWalletCharge(&data)
	if chargeErr != nil {
		log.Fatal(chargeErr)
	}
	fmt.Printf("created e-wallet charge: %+v\n", charge)
}

func ExampleGetEWalletChargeStatus() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := ewallet.GetEWalletChargeStatusParams{
		ChargeID: "ewc_f3925450-5c54-4777-98c1-fcf22b0d1e1c",
	}

	charge, chargeErr := ewallet.GetEWalletChargeStatus(&data)
	if chargeErr != nil {
		log.Fatal(chargeErr)
	}
	fmt.Printf("retrieved e-wallet charge: %+v\n", charge)
}

package main

import (
	"fmt"
	"log"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/account"
)

func accountTest() {
	data := account.CreateParams{
		Email: "example@mail.com",
		Type:  xendit.OWNED,
		PublicProfile: xendit.PublicProfile{
			BusinessName: "John Doe",
		},
	}
	resp, err := account.Create(&data)
	if err != nil {
		log.Panic(err)

	}
	resp, err = account.Get(&account.GetParams{
		ID: resp.ID,
	})
	if err != nil {
		log.Panic(err)
	}
	if resp.Email != data.Email {
		log.Panic("Email is not equal")
	}
	if resp.Type != data.Type {
		log.Panic("Type is not equal")
	}
	fmt.Println("Invoice integration tests done!")
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/account"
)

func accountTest() {
	rand.Seed(time.Now().UnixNano())
	email := strconv.Itoa(rand.Intn(100000000)) + "@example.com" //Generate more or less random email
	// fmt.Println(email)
	data := account.CreateParams{
		Email: email,
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
	fmt.Println("Account integration tests done!")
}

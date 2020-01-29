package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/balance"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	getData := balance.GetParams{
		AccountType: "CASH",
	}

	resp, err := balance.Get(&getData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("balance: %+v\n", resp)
}

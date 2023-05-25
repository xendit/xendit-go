package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/report"
	"log"
	"os"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}

	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	generateReportData := report.GenerateReportParams{
		Type: "BALANCE_HISTORY", // "BALANCE_HISTORY", "TRANSACTIONS", "UPCOMING_TRANSACTIONS"
		Filter: report.Filter{
			From: "2020-01-01T00:00:00.000Z",
			To:   "2020-12-31T23:59:59.999Z",
		},
	}

	resp, err := report.GenerateReport(&generateReportData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("generated report: %+v\n", resp)

	getReportData := report.GetReportParams{
		ReportID: resp.ID,
	}

	resp, err = report.GetReport(&getReportData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved report: %+v\n", resp)
}

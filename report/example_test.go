package report

import (
	"github.com/xendit/xendit-go"
	"log"
)

func ExampleGenerateReport() {
	xendit.Opt.SecretKey = "examplesecretkey"

	generateReport := GenerateReportParams{
		Type: "BALANCE_HISTORY", // "BALANCE_HISTORY", "TRANSACTIONS", "UPCOMING_TRANSACTIONS"
		Filter: Filter{
			From: "2020-01-01T00:00:00.000Z",
			To:   "2020-12-31T23:59:59.999Z",
		},
	}

	resp, err := GenerateReport(&generateReport)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}

	log.Printf("generated report: %+v\n", resp)
}

func ExampleGetReport() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getReport := GetReportParams{
		ReportID: "report_5c1b34a2-6ceb-4c24-aba9-c836bac82b28",
	}

	resp, err := GetReport(&getReport)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}

	log.Printf("retrieved report: %+v\n", resp)
}

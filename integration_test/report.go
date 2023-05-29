package main

import (
	"fmt"
	"github.com/xendit/xendit-go/report"
	"log"
)

func reportTest() {
	respReport, err := report.GenerateReport(&report.GenerateReportParams{
		Type: "BALANCE_HISTORY", // "BALANCE_HISTORY", "TRANSACTIONS", "UPCOMING_TRANSACTIONS"
		Filter: report.Filter{
			From: "2020-01-01T00:00:00.000Z",
			To:   "2020-12-31T23:59:59.999Z",
		},
	})
	if err != nil {
		log.Panic(err)
	}

	_, err = report.GetReport(&report.GetReportParams{
		ReportID: respReport.ID,
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("report integration tests done!")
}

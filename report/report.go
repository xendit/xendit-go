package report

import (
	"context"
	"github.com/xendit/xendit-go"
)

func GenerateReport(data *GenerateReportParams) (*xendit.Report, *xendit.Error) {
	return GenerateReportWithContext(context.Background(), data)
}

func GenerateReportWithContext(ctx context.Context, data *GenerateReportParams) (*xendit.Report, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GenerateReportWithContext(ctx, data)
}

func GetReport(data *GetReportParams) (*xendit.Report, *xendit.Error) {
	return GetReportWithContext(context.Background(), data)
}

func GetReportWithContext(ctx context.Context, data *GetReportParams) (*xendit.Report, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetReportWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}

package report

import (
	"context"
	"fmt"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
	"net/http"
)

type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// GenerateReport creates a report
func (c *Client) GenerateReport(data *GenerateReportParams) (*xendit.Report, *xendit.Error) {
	return c.GenerateReportWithContext(context.Background(), data)
}

// GenerateReportWithContext creates a report with context
func (c *Client) GenerateReportWithContext(ctx context.Context, data *GenerateReportParams) (*xendit.Report, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Report{}

	header := http.Header{}
	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/reports", c.Opt.XenditURL),
		c.Opt.SecretKey,
		header,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetReport gets a report
func (c *Client) GetReport(data *GetReportParams) (*xendit.Report, *xendit.Error) {
	return c.GetReportWithContext(context.Background(), data)
}

// GetReportWithContext gets a report with context
func (c *Client) GetReportWithContext(ctx context.Context, data *GetReportParams) (*xendit.Report, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Report{}

	header := http.Header{}
	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/reports/%s", c.Opt.XenditURL, data.ReportID),
		c.Opt.SecretKey,
		header,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

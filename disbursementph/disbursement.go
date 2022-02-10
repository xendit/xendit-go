package disbursementph

import (
	"context"

	"github.com/xendit/xendit-go"
)

// Create creates new disbursement
func Create(data *CreateParams) (*xendit.DisbursementPh, *xendit.Error) {
	return CreateWithContext(context.Background(), data)
}

// CreateWithContext creates new disbursement with context
func CreateWithContext(ctx context.Context, data *CreateParams) (*xendit.DisbursementPh, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateWithContext(ctx, data)
}

// GetByID gets a disbursement
func GetByID(data *GetByIDParams) (*xendit.DisbursementPh, *xendit.Error) {
	return GetByIDWithContext(context.Background(), data)
}

// GetByIDWithContext gets a disbursement with context
func GetByIDWithContext(ctx context.Context, data *GetByIDParams) (*xendit.DisbursementPh, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetByIDWithContext(ctx, data)
}

// GetByReferenceID gets a disbursement
func GetByReferenceID(data *GetByReferenceIDParams) ([]xendit.DisbursementPh, *xendit.Error) {
	return GetByReferenceIDWithContext(context.Background(), data)
}

// GetByExternalIDWithContext gets a disbursement with context
func GetByReferenceIDWithContext(ctx context.Context, data *GetByReferenceIDParams) ([]xendit.DisbursementPh, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetByReferenceIDWithContext(ctx, data)
}

// GetDisbursementChannels gets available disbursement banks
func GetDisbursementChannels() ([]xendit.DisbursementChannel, *xendit.Error) {
	return GetDisbursementChannelsWithContext(context.Background())
}

// GetDisbursementChannelsWithContext gets available disbursement banks with context
func GetDisbursementChannelsWithContext(ctx context.Context) ([]xendit.DisbursementChannel, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetDisbursementChannelsWithContext(ctx)
}

// GetDisbursementChannelsByChannelCategory gets available disbursement banks
func GetDisbursementChannelsByChannelCategory(data *GetByChannelCategoryParams) ([]xendit.DisbursementChannel, *xendit.Error) {
	return GetDisbursementChannelsByChannelCategoryWithContext(context.Background(), data)
}

// GetDisbursementChannelsByChannelCategoryWithContext gets available disbursement banks with context
func GetDisbursementChannelsByChannelCategoryWithContext(ctx context.Context, data *GetByChannelCategoryParams) ([]xendit.DisbursementChannel, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetDisbursementChannelsByChannelCategoryWithContext(ctx, data)
}

// GetDisbursementChannelsByChannelCode gets available disbursement banks
func GetDisbursementChannelsByChannelCode(data *GetByChannelCodeParams) ([]xendit.DisbursementChannel, *xendit.Error) {
	return GetDisbursementChannelsByChannelCodeWithContext(context.Background(), data)
}

// GetDisbursementChannelsByChannelCodeWithContext gets available disbursement banks with context
func GetDisbursementChannelsByChannelCodeWithContext(ctx context.Context, data *GetByChannelCodeParams) ([]xendit.DisbursementChannel, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetDisbursementChannelsByChannelCodeWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}

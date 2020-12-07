package promotion

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreatePromotion creates new promotion
func CreatePromotion(data *CreatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	return CreatePromotionWithContext(context.Background(), data)
}

// CreatePromotionWithContext creates new promotion with context
func CreatePromotionWithContext(ctx context.Context, data *CreatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreatePromotionWithContext(ctx, data)
}

// GetPromotions gets promotions
func GetPromotions(data *GetPromotionsParams) ([]xendit.Promotion, *xendit.Error) {
	return GetPromotionsWithContext(context.Background(), data)
}

// GetPromotionsWithContext gets promotions with context
func GetPromotionsWithContext(ctx context.Context, data *GetPromotionsParams) ([]xendit.Promotion, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPromotionsWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}

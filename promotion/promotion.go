package promotion

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreatePromotion creates new promotion.
func CreatePromotion(data *CreatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	return CreatePromotionWithContext(context.Background(), data)
}

// CreatePromotionWithContext creates new promotion with context.
func CreatePromotionWithContext(ctx context.Context, data *CreatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreatePromotionWithContext(ctx, data)
}

// GetPromotions gets promotions.
func GetPromotions(data *GetPromotionsParams) ([]xendit.Promotion, *xendit.Error) {
	return GetPromotionsWithContext(context.Background(), data)
}

// GetPromotionsWithContext gets promotions with context.
func GetPromotionsWithContext(ctx context.Context, data *GetPromotionsParams) ([]xendit.Promotion, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPromotionsWithContext(ctx, data)
}

// GetPromotionsCalculation gets promotions calculation.
func GetPromotionsCalculation(data *GetPromotionsCalculationParams) (*xendit.PromotionCalculation, *xendit.Error) {
	return GetPromotionsCalculationWithContext(context.Background(), data)
}

// GetPromotionsCalculationWithContext gets promotions calculation with context.
func GetPromotionsCalculationWithContext(ctx context.Context, data *GetPromotionsCalculationParams) (*xendit.PromotionCalculation, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPromotionsCalculationWithContext(ctx, data)
}

// UpdatePromotion updates a promotion.
func UpdatePromotion(data *UpdatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	return UpdatePromotionWithContext(context.Background(), data)
}

// UpdatePromotionWithContext updates a promotion with context.
func UpdatePromotionWithContext(ctx context.Context, data *UpdatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.UpdatePromotionWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}

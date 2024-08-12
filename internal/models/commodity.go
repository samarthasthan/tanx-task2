package models

type Commodity struct {
	Name        string  `json:"item_name" validate:"required"`
	Description string  `json:"item_description" validate:"required"`
	Price       float32 `json:"quote_price_per_month" validate:"required"`
	Category    string  `json:"item_category" validate:"required"`
}

type CommodityOut struct {
	CommodityId string  `json:"commodity_id" validate:"required"`
	Price       float32 `json:"quote_price_per_month" validate:"required"`
	CreatedAt   string  `json:"created_at" validate:"required"`
}

type CommodityResponse struct {
	CommodityId string  `json:"commodity_id" validate:"required"`
	Price       string `json:"quote_price_per_month" validate:"required"`
	CreatedAt   string  `json:"created_at" validate:"required"`
	Category    string  `json:"item_category" validate:"required"`
}

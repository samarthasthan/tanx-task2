package models

type Bid struct {
	CommodityId string `json:"commodity_id" validate:"required"`
	Price       string `json:"bid_price_month" validate:"required"`
	Duration    int    `json:"rental_duration" validate:"required"`
}

type BidOut struct {
	BidId       string `json:"bid_id" validate:"required"`
	CommodityId string `json:"commodity_id" validate:"required"`
	CreatedAt   string `json:"created_at" validate:"required"`
}

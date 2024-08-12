package models

type Bid struct {
	CommodityId        string `json:"commodity_id" validate:"required"`
	Price 			string `json:"bid_price_month" validate:"required"`
	Duration 		string `json:"rental_duration" validate:"required"`
}
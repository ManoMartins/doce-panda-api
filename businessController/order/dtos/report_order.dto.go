package dtos

import "time"

type InputReportOrderDto struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type Order struct {
	CategoryName string    `json:"categoryName"`
	PriceInCents int       `json:"priceInCents"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Series struct {
	CategoryID   string  `json:"categoryId"`
	CategoryName string  `json:"categoryName"`
	Orders       []Order `json:"orders"`
}

type OutputReportOrderDto struct {
	Series    []Series  `json:"series"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

package types

import (
	"time"
)

type ProductReport struct {
	EAN             string  `json:"ean" validate:"required,min=8,max=13"`
	ExternalId      string  `json:"externalId"`
	Name            string  `json:"name" validate:"required"`
	Description     string  `json:"description"`
	Price           float32 `json:"price" validate:"required,gte=0"`
	DiscountedPrice float32 `json:"discountedPrice" validate:"gte=0"`
	Brand           string  `json:"brand" validate:"required"`
	Image           string  `json:"image"`
	Type            string  `json:"type" validate:"required"`
}

type ProductReportMessage struct {
	EAN             string    `json:"ean"`
	ExternalId      string    `json:"externalId"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           float32   `json:"price"`
	DiscountedPrice float32   `json:"discountedPrice"`
	Brand           string    `json:"brand"`
	Image           string    `json:"image"`
	Type            string    `json:"type"`
	CreationTime    time.Time `json:"publishTime"`
}

func MapSaveRequestToQueueDTO(report *ProductReport) ProductReportMessage {
	return ProductReportMessage{
		EAN:             report.EAN,
		ExternalId:      report.ExternalId,
		Name:            report.Name,
		Description:     report.Description,
		Price:           report.Price,
		DiscountedPrice: report.DiscountedPrice,
		Brand:           report.Brand,
		Image:           report.Image,
		Type:            report.Type,
		CreationTime:    time.Now(),
	}
}

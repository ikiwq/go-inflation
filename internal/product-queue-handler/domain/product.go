package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID           int       `db:"id"`
	EAN          string    `db:"ean"`
	ExternalId   string    `db:"external_id"`
	Name         string    `db:"name"`
	Description  string    `db:"Description"`
	Brand        string    `db:"brand"`
	Image        string    `db:"image"`
	Type         string    `db:"type"`
	CreationTime time.Time `db:"creation_time"`
	SaveTime     time.Time `db:"save_time"`
}

type ProductSqlRepository interface {
	Save(context.Context, Product) (Product, error)
	FindByEan(context.Context, string) (Product, error)
}

type ProductPriceHistory struct {
	ID              int       `db:"id"`
	ProductId       int       `db:"product_id"`
	Price           float32   `db:"price"`
	DiscountedPrice float32   `db:"discounted_price"`
	RegisteredAt    time.Time `db:"registered_at"`
	SaveTime        time.Time `db:"save_time"`
}

type ProductPriceHistorySqlRepository interface {
	Save(context.Context, ProductPriceHistory) (ProductPriceHistory, error)
}

type ProductDocument struct {
	EAN             string    `json:"ean"`
	ExternalId      string    `json:"externalId"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description"`
	Price           float32   `json:"price"`
	DiscountedPrice float32   `json:"discountedPrice"`
	Brand           string    `json:"brand"`
	Image           string    `json:"image"`
	Type            string    `json:"type"`
	CreationTime    time.Time `json:"creationTime"`
	SaveTime        time.Time `json:"saveTime"`
}

type ProductDocumentMongoRepository interface {
	Save(context.Context, ProductDocument) (*mongo.InsertOneResult, error)
}

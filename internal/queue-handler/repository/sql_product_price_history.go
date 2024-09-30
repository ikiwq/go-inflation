package repository

import (
	"context"

	"github.com/ikiwq/go-inflation/internal/queue-handler/domain"
	"github.com/jmoiron/sqlx"
)

type productPriceHistoryRepository struct {
	dbClient *sqlx.DB
}

func NewProductPriceHistoryRepository(dbClient *sqlx.DB) domain.ProductPriceHistorySqlRepository {
	return &productPriceHistoryRepository{dbClient: dbClient}
}

func (r *productPriceHistoryRepository) Save(ctx context.Context, productPriceHistory *domain.ProductPriceHistory) (*domain.ProductPriceHistory, error) {
	query := `
		INSERT INTO products_price_history VALUES(
			:product_id, :price, :discounted_price, :registered_at, :save_time
		)
	`

	res, err := r.dbClient.NamedExecContext(ctx, query, productPriceHistory)
	if err != nil {
		return nil, err
	}

	lastInsertedId, _ := res.LastInsertId()
	productPriceHistory.ID = int(lastInsertedId)

	return productPriceHistory, nil
}

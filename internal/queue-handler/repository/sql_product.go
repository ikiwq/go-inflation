package repository

import (
	"github.com/ikiwq/go-inflation/internal/queue-handler/domain"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type productRepository struct {
	dbClient *sqlx.DB
}

func NewSqlProductRepository(dbClient *sqlx.DB) domain.ProductSqlRepository {
	return &productRepository{dbClient: dbClient}
}

func (r *productRepository) Save(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `
		INSERT INTO products 
		VALUES(
			:ean, :external_id, :name, :description, 
			:brand, :image, :type, :creation_time, :save_time
		)
	`
	res, err := r.dbClient.NamedExecContext(ctx, query, product)
	if err != nil {
		return domain.Product{}, err
	}

	lastInsertedId, _ := res.LastInsertId()
	product.ID = int(lastInsertedId)

	return product, nil
}

func (r *productRepository) FindByEan(ctx context.Context, ean string) (domain.Product, error) {
	var product domain.Product

	err := r.dbClient.GetContext(ctx, &product, "SELECT * FROM products WHERE ean = ?", ean)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

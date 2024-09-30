package repository

import (
	"github.com/ikiwq/go-inflation/internal/queue-handler/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type productDocumentRepository struct {
	mongoCollection *mongo.Collection
}

func NewMongoProductRepository(mongoCollection *mongo.Collection) domain.ProductDocumentMongoRepository {
	return &productDocumentRepository{mongoCollection: mongoCollection}
}

func (r *productDocumentRepository) Save(ctx context.Context, productDocument *domain.ProductDocument) (*mongo.InsertOneResult, error) {
	res, err := r.mongoCollection.InsertOne(ctx, productDocument)
	return res, err
}

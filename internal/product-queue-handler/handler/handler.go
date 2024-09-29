package handler

import (
	"time"

	"github.com/ikiwq/go-inflation/internal/product-queue-handler/config"
	"github.com/ikiwq/go-inflation/internal/product-queue-handler/domain"
	"github.com/ikiwq/go-inflation/internal/product-queue-handler/repository"
	"github.com/ikiwq/go-inflation/pkg/db"
	"github.com/ikiwq/go-inflation/pkg/queue"
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	productKafkaReader *kafka.Reader

	sqlClient     *sqlx.DB
	mongoDatabase *mongo.Database

	productDocumentRepository domain.ProductDocumentMongoRepository

	productRepository             domain.ProductSqlRepository
	productPriceHistoryRepository domain.ProductPriceHistorySqlRepository
}

func NewHandler(config *config.ProductQueueHandlerConfig) *handler {
	productKafkaReader := queue.InitReader(
		[]string{config.Kafka.Address},
		config.Kafka.ProductCreationRequestConfig.Topic,
		100*time.Millisecond,
		10e3,
		10e6,
	)

	sqlClient := db.InitSqlClient(
		config.SQL.DriverName,
		config.SQL.Username,
		config.SQL.Password,
		config.SQL.Address,
		config.SQL.DbName,
	)

	mongoDatabase := db.InitMongoDatabase(
		config.MongoDB.Username,
		config.MongoDB.Password,
		config.MongoDB.Address,
		config.MongoDB.DbName,
	)

	productCollection := mongoDatabase.Collection("ProductsCollection")
	mongoProductRepository := repository.NewMongoProductRepository(productCollection)

	

	return &handler{
		productKafkaReader:        productKafkaReader,

		sqlClient:                 sqlClient,
		mongoDatabase:             mongoDatabase,

		productDocumentRepository: mongoProductRepository,

		productRepository: ,
	}
}

func (h *handler) Start() {
	h.listenForProductMessages()
}

func (h *handler) Exit() {
	h.productKafkaReader.Close()
	h.sqlClient.Close()
}

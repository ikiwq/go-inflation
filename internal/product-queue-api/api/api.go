package api

import (
	"fmt"
	"net/http"

	"github.com/ikiwq/go-inflation/internal/product-queue-api/config"
	"github.com/ikiwq/go-inflation/pkg/queue"
	"github.com/segmentio/kafka-go"
)

type api struct {
	apiAddress string
	apiPort    string
	httpClient *http.Client

	productKafkaConn *kafka.Conn
}

func NewApi(config config.ProductQueueApiConfig) *api {
	httpClient := &http.Client{}

	productKafkaConn := queue.InitConnection(
		config.Kafka.NetworkType,
		config.Kafka.Address,
		config.Kafka.ProductCreationRequestConfig.Topic,
		config.Kafka.ProductCreationRequestConfig.Partition,
	)

	return &api{
		apiAddress: config.Server.Host,
		apiPort:    config.Server.Port,
		httpClient: httpClient,

		productKafkaConn: productKafkaConn,
	}
}

func (a *api) buildRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /api/v1/products", a.saveProduct)

	return r
}

func (a *api) Start() {
	r := a.buildRoutes()

	listenAddress := fmt.Sprintf("%s:%s", a.apiAddress, a.apiPort)
	http.ListenAndServe(listenAddress, r)
}

func (a *api) Exit() {
	a.httpClient.CloseIdleConnections()
	a.productKafkaConn.Close()
}

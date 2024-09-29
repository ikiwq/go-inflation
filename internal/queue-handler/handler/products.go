package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ikiwq/go-inflation/internal/queue-handler/domain"
	"github.com/ikiwq/go-inflation/pkg/types"
)

func (h *handler) listenForProductMessages() {
	for {
		m, err := h.productKafkaReader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("error while reading message: ", err)
			break
		}

		fmt.Printf("received message: %s \n", string(m.Value))

		var productMessage types.ProductReportMessage
		err = json.Unmarshal(m.Value, &productMessage)
		if err != nil {
			fmt.Printf("error while unmarshaling message %s \n", string(m.Value))
		}

		productDocument := domain.ProductDocument{
			EAN:             productMessage.EAN,
			ExternalId:      productMessage.ExternalId,
			Name:            productMessage.Name,
			Description:     productMessage.Description,
			DiscountedPrice: productMessage.DiscountedPrice,
			Brand:           productMessage.Brand,
			Image:           productMessage.Image,
			Type:            productMessage.Type,
			CreationTime:    productMessage.CreationTime,
			SaveTime:        time.Now(),
		}

		_, err = h.productDocumentRepository.Save(context.TODO(), productDocument)
		if err != nil {
			fmt.Printf("error while saving document %s", err)
		}
	}
}

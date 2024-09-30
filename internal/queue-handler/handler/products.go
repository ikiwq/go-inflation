package handler

import (
	"context"
	"encoding/json"
	"fmt"

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

		err = h.saveProductMessageAsDocument(&productMessage)
		if err != nil {
			fmt.Printf("error while saving document %s", err)
		}

		err = h.getOrSaveProductInformation(&productMessage)
		if err != nil {
			fmt.Printf("error while saving product %s", err)
		}
	}
}

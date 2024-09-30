package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/ikiwq/go-inflation/pkg/types"
	"github.com/ikiwq/go-inflation/pkg/utils"
)

func (a *api) saveProduct(w http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var saveRequest types.ProductReport
	err = json.Unmarshal(body, &saveRequest)
	if err != nil {
		http.Error(w, "Error while parsing JSON", http.StatusBadRequest)
		return
	}

	validate := validator.New()

	err = validate.Struct(saveRequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	queueDTO := types.MapSaveRequestToQueueDTO(&saveRequest)
	dtoBytes, err := json.Marshal(queueDTO)
	if err != nil {
		http.Error(w, "Unexpected error happened while handling save", http.StatusInternalServerError)
	}

	_, err = a.productKafkaConn.Write(dtoBytes)
	if err != nil {
		fmt.Print("Error while saving message", err)
		http.Error(w, "Unexpected error happened while handling publishing to the queue", http.StatusInternalServerError)
	}

	utils.WriteJSON(w, http.StatusCreated, queueDTO)
}

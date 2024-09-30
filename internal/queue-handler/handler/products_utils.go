package handler

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ikiwq/go-inflation/internal/queue-handler/domain"
	"github.com/ikiwq/go-inflation/pkg/types"
)

func (h *handler) saveProductMessageAsDocument(productMessage *types.ProductReportMessage) error {
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

	_, err := h.productDocumentRepository.Save(context.Background(), &productDocument)

	return err
}

func (h *handler) getOrSaveProductInformation(productMessage *types.ProductReportMessage) error {
	prevProduct, err := h.productRepository.FindByEan(context.Background(), productMessage.EAN)
	updatedProduct := &domain.Product{
		EAN:          productMessage.EAN,
		ExternalId:   productMessage.ExternalId,
		Name:         productMessage.Name,
		Description:  productMessage.Description,
		Brand:        productMessage.Brand,
		Image:        productMessage.Image,
		Type:         productMessage.Type,
		CreationTime: productMessage.CreationTime,
		SaveTime:     time.Now(),
	}

	if err == nil {
		updatedProduct.ID = prevProduct.ID
		h.productRepository.Update(context.Background(), updatedProduct)
		return nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	_, saveErr := h.productRepository.Save(context.Background(), updatedProduct)
	if saveErr != nil {
		return err
	}
	return nil
}

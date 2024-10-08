package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
	"strconv"

	graphModel "github.com/ngocxxu/grocery-store-svelte-be/graph/model"
	internalModel "github.com/ngocxxu/grocery-store-svelte-be/internal/model"
	"github.com/ngocxxu/grocery-store-svelte-be/pkg/utils"
)

// CreateProduct is the resolver for the createProduct field.
func (r *Resolver) CreateProduct(ctx context.Context, input *graphModel.ProductInput) (*graphModel.Product, error) {
	// Convert graph type to internal type
	internalProduct := &internalModel.Product{
		Name:          input.Name,
		Description:   input.Description,
		Type:          input.Type,
		Sku:           input.Sku,
		Status:        input.Status,
		Price:         input.Price,
		Discount:      input.Discount,
		Rating:        uint(input.Rating),
		Quantity:      uint(input.Quantity),
		WeightOptions: make([]internalModel.WeightOption, len(input.WeightOptions)),
	}

	for i, optInput := range input.WeightOptions {
		unitID, err := strconv.ParseUint(optInput.UnitID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid unit ID: %v", err)
		}
		internalProduct.WeightOptions[i] = internalModel.WeightOption{
			Weight: optInput.Weight,
			UnitID: uint(unitID),
		}
	}

	product, err := r.ProductService.CreateProduct(internalProduct)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToGraphProduct(product), nil
}

// Products is the resolver for the products field.
func (r *Resolver) Products(ctx context.Context) ([]*graphModel.Product, error) {
	products, err := r.ProductService.GetProducts()
	if err != nil {
		return nil, err
	}
	var result []*graphModel.Product
	for _, product := range products {
		result = append(result, utils.ConvertToGraphProduct(product))
	}
	return result, nil
}

// Product is the resolver for the product field.
func (r *Resolver) Product(ctx context.Context, id string) (*graphModel.Product, error) {
	product, err := r.ProductService.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return utils.ConvertToGraphProduct(product), nil
}

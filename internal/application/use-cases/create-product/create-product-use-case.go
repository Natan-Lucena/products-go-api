package useCases

import (
	"time"

	"github.com/Natan-Lucena/products-api/internal/application/clients"
	appErrors "github.com/Natan-Lucena/products-api/internal/application/errors"
	domain "github.com/Natan-Lucena/products-api/internal/domain/entities"
	"github.com/Natan-Lucena/products-api/internal/domain/repositories"
	"github.com/google/uuid"
)

type CreateProductUseCase struct {
	productRepository repositories.ProductRepository
	userAppClient clients.UserAppClient
}

func NewCreateProductUseCase(productRepository repositories.ProductRepository, 
	userAppClient clients.UserAppClient) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepository: productRepository,
		userAppClient: userAppClient,
	}
}

func (c *CreateProductUseCase) Execute(userId uuid.UUID, name string, price float64, quantity int) (*domain.Product, error) {
	user, err := c.userAppClient.GetUserByID(userId)
	if err != nil {
		return nil, err
	}
	if user.Role() != "ADMIN" {
		return nil, appErrors.ErrUnauthorized
	}
	product, err := domain.NewProduct(uuid.Nil, user.ID(), name, price, quantity, time.Time{})
	if err != nil {
		return nil, err
	}
	response , err := c.productRepository.Save(product)
	if err != nil {
		return nil, err
	}
	return response, nil
}
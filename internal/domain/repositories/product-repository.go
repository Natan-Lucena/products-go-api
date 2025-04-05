package repositories

import domain "github.com/Natan-Lucena/products-api/internal/domain/entities"

type ProductRepository interface {
	Save(product *domain.Product) (*domain.Product, error)
}
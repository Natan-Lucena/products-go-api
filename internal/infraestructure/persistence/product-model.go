package models

import (
	"time"

	domain "github.com/Natan-Lucena/products-api/internal/domain/entities"
	"github.com/google/uuid"
)

type ProductModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

func ToDomainProduct(m *ProductModel) (*domain.Product, error) {
	return domain.NewProduct(
		m.ID,
		m.UserID,
		m.Name,
		m.Price,
		m.Quantity,
		m.CreatedAt,
	)
}

func FromDomainProduct(p *domain.Product) *ProductModel {
	return &ProductModel{
		ID:        p.ID(),
		UserID:    p.UserID(),
		Name:      p.Name(),
		Price:     p.Price(),
		Quantity:  p.Quantity(),
		CreatedAt: p.CreatedAt(),
	}
}

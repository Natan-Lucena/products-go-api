package models

import (
	"time"

	domain "github.com/Natan-Lucena/products-api/internal/domain/entities"
	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

// Converte para a entidade
func ToDomain(m *UserModel) *domain.User {
	user, err := domain.NewUser(
		m.ID,
		m.Name,
		m.Email,
		m.Password,
		domain.Role(m.Role),
		m.CreatedAt,
	)
	if err != nil {
		return nil
	}
	return user
}

func FromDomain(u *domain.User) *UserModel {
	return &UserModel{
		ID:        u.ID(),
		Name:      u.Name(),
		Email:     u.Email(),
		Password:  u.Password(),
		Role:      string(u.Role()),
		CreatedAt: u.CreatedAt(),
	}
}

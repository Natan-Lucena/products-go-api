package clients

import (
	domain "github.com/Natan-Lucena/products-api/internal/domain/entities"
	"github.com/google/uuid"
)

type UserAppClient interface {
	GetUserByID(userID uuid.UUID) (*domain.User, error)
}

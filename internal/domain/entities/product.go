package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	id        uuid.UUID
	userId    uuid.UUID
	name      string
	price     float64
	quantity  int
	createdAt time.Time
}

func NewProduct(
	id uuid.UUID,
	userId uuid.UUID,
	name string,
	price float64,
	quantity int,
	createdAt time.Time,
) (*Product, error) {
	if name == "" || price < 0 || quantity < 0 || userId == uuid.Nil {
		return nil, errors.New("INVALID_ARGUMENTS")
	}
	if id == uuid.Nil {
		id = uuid.New()
	}
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	return &Product{
		id:        id,
		userId:    userId,
		name:      name,
		price:     price,
		quantity:  quantity,
		createdAt: createdAt,
	}, nil
}

func (p *Product) ID() uuid.UUID         { return p.id }
func (p *Product) UserID() uuid.UUID     { return p.userId }
func (p *Product) Name() string          { return p.name }
func (p *Product) Price() float64        { return p.price }
func (p *Product) Quantity() int         { return p.quantity }
func (p *Product) CreatedAt() time.Time  { return p.createdAt }

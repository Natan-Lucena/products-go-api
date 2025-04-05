package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	name      string
	email     string
	role      Role
	createdAt time.Time
}

func NewUser(
	id uuid.UUID,
	name, email string,
	role Role,
	createdAt time.Time,
) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("INVALID_ARGUMENTS")
	}
	if !IsValidRole(role) {
		return nil, errors.New("INVALID_ROLE")
	}

	if id == uuid.Nil {
		id = uuid.New()
	}
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	return &User{
		id:        id,
		name:      name,
		email:     email,
		role:      role,
		createdAt: createdAt,
	}, nil
}
func (u *User) ID() uuid.UUID         { return u.id }
func (u *User) Name() string          { return u.name }
func (u *User) Email() string         { return u.email }
func (u *User) Role() Role            { return u.role }
func (u *User) CreatedAt() time.Time  { return u.createdAt }
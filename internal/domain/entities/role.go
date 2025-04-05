package domain

type Role string

const (
	AdminRole Role = "ADMIN"
	UserRole  Role = "USER"
)

func IsValidRole(role Role) bool {
	return role == AdminRole || role == UserRole
}

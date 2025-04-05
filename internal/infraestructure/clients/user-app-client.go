package clientImpl

import (
	"context"
	"time"

	"github.com/Natan-Lucena/products-api/internal/application/clients"
	domain "github.com/Natan-Lucena/products-api/internal/domain/entities"
	user "github.com/Natan-Lucena/products-api/internal/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type userGRPCClient struct {
	client user.UserServiceClient
}

func NewUserGRPCClient(conn *grpc.ClientConn) clients.UserAppClient {
	return &userGRPCClient{
		client: user.NewUserServiceClient(conn),
	}
}

func (c *userGRPCClient) GetUserByID(userID uuid.UUID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &user.GetUserRequest{
		UserId: userID.String(),
	}

	res, err := c.client.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	userData := res.GetUser()
	role := domain.Role(userData.GetRole()) 

	createdAt, err := time.Parse(time.RFC3339, userData.GetCreatedAt())
	if err != nil {
		return nil, err
	}

	return domain.NewUser(
		uuid.MustParse(userData.GetId()),
		userData.GetName(),
		userData.GetEmail(),
		role,
		createdAt,
	)
}

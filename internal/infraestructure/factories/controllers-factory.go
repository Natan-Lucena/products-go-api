package factories

import (
	"context"
	"time"

	useCases "github.com/Natan-Lucena/products-api/internal/application/use-cases/create-product"
	clientImpl "github.com/Natan-Lucena/products-api/internal/infraestructure/clients"
	repositoriesImpl "github.com/Natan-Lucena/products-api/internal/infraestructure/repositories"
	"google.golang.org/grpc"
)

func GetCreateProductController() *useCases.CreateProductController {
	productRepository := repositoriesImpl.NewProductRepository()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, _ := grpc.DialContext(ctx, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	userAppClient := clientImpl.NewUserGRPCClient(conn)
	createProductUseCase := useCases.NewCreateProductUseCase(productRepository, userAppClient)
	return useCases.NewCreateProductController(*createProductUseCase)

}
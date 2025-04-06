package repositoriesImpl

import (
	domain "github.com/Natan-Lucena/products-api/internal/domain/entities"
	"github.com/Natan-Lucena/products-api/internal/domain/repositories"
	models "github.com/Natan-Lucena/products-api/internal/infraestructure/persistence"
	gormPkg "github.com/Natan-Lucena/products-api/pkg/gorm"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository() repositories.ProductRepository {
	db, _ := gormPkg.InitDB()
	return &productRepositoryImpl{
		db: db,
	}
}

func (r *productRepositoryImpl) Save(product *domain.Product) (*domain.Product, error) {

	productModel := models.FromDomainProduct(product)

	if err := r.db.Save(productModel).Error; err != nil {
		return nil, err
	}

	return product, nil
}
	

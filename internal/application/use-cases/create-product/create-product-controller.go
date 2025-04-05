package useCases

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateProductController struct {
	productUseCase CreateProductUseCase
}

func NewCreateProductController(productUseCase CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		productUseCase: productUseCase,
	}
}

func (controller *CreateProductController) Handle(ctx *gin.Context) {
	var request struct {
		Name     string  `json:"name" binding:"required"`
		Price    float64 `json:"price" binding:"required"`
		Quantity int     `json:"quantity" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	userIdStr := ctx.GetString("userId")
	if userIdStr == "" {
		ctx.JSON(401, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "userId inválido"})
		return
	}

	product, err := controller.productUseCase.Execute(userId, request.Name, request.Price, request.Quantity)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao criar produto", "details": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{
		"name":     product.Name(),
		"price":    product.Price(),
		"quantity": product.Quantity(),
		"createdAt": product.CreatedAt(),
	})
}
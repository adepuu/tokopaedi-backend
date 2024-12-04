package http

import (
	"github.com/adepuu/tokopaedi-backend/internal/model"
	"github.com/adepuu/tokopaedi-backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	Usecase usecase.ProductUsecase
	Log     *logrus.Logger
}

func NewProductController(uc *usecase.ProductUsecase, log *logrus.Logger) *ProductController {
	return &ProductController{
		Usecase: *uc,
		Log:     log,
	}
}

// CreateProduct implements ProductController.
func (p *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).JSON(model.WebResponse[any]{
		Message: "create product not implemented",
		Success: false,
	})
}

// GetProductByID implements ProductController.
func (p *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).JSON(model.WebResponse[any]{
		Message: "get product by id not implemented",
		Success: false,
	})
}

// UpdateProduct implements ProductController.
func (p *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).JSON(model.WebResponse[any]{
		Message: "update product not implemented",
		Success: false,
	})
}

// DeleteProduct implements ProductController.
func (p *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).JSON(model.WebResponse[any]{
		Message: "delete product not implemented",
		Success: false,
	})
}

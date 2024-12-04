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
	request := new(model.CreateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}
	resp, err := p.Usecase.CreateProduct(ctx.Context(), request)
	if err != nil {
		p.Log.WithError(err).Error("failed to create product")
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(model.WebResponse[*model.CreateProductResponse]{Data: resp})
	}
	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CreateProductResponse]{Data: resp})
}

// DeleteProduct implements ProductController.
func (p *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// GetProductByID implements ProductController.
func (p *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateProduct implements ProductController.
func (p *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

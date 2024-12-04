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
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}
	resp, err := p.Usecase.CreateProduct(ctx.Context(), request)
	if err != nil {
		p.Log.WithError(err).Error("failed to create product")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.CreateProductResponse]{
			Data:    resp,
			Success: false,
			Message: "failed to create product",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CreateProductResponse]{
		Data:    resp,
		Success: true,
		Message: "product created successfully",
	})
}

// GetProductByID implements ProductController.
func (p *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid product id format",
			Success: false,
		})
	}

	product, err := p.Usecase.GetProductByID(ctx.Context(), int64(id))
	if err != nil {
		if err.Error() == "product not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
				Message: "product not found",
				Success: false,
			})
		}

		p.Log.WithError(err).Error("failed to get product")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve product details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CreateProductResponse]{
		Data:    product,
		Success: true,
		Message: "product retrieved successfully",
	})
}

// UpdateProduct implements ProductController.
func (p *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid product id format",
			Success: false,
		})
	}

	request := new(model.UpdateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}

	product, err := p.Usecase.UpdateProduct(ctx.Context(), request, int64(id))
	if err != nil {
		p.Log.WithError(err).Error("failed to update product")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to update product details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CreateProductResponse]{
		Data:    product,
		Success: true,
		Message: "product updated successfully",
	})
}

// DeleteProduct implements ProductController.
func (p *ProductController) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid product id format",
			Success: false,
		})
	}

	resp, err := p.Usecase.DeleteProduct(ctx.Context(), int64(id))
	if err != nil {
		p.Log.WithError(err).Error("failed to delete product")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to delete product",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.DeleteProductResponse]{
		Data:    resp,
		Success: true,
		Message: "product deleted successfully",
	})
}

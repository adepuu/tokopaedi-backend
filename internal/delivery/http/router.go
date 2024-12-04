package http

import "github.com/gofiber/fiber/v2"

type Router struct {
	App               *fiber.App
	ProductController *ProductController
	AuthMiddleware    fiber.Handler
}

type router interface {
	Setup()
	registerPublicEndpoints()
	registerPrivateEndpoints()
}

func NewRouter(app *fiber.App, productController *ProductController, authMiddleware fiber.Handler) router {
	return &Router{
		App:               app,
		ProductController: productController,
		AuthMiddleware:    authMiddleware,
	}
}

// Setup implements router.
func (r *Router) Setup() {
	r.registerPublicEndpoints()
	r.registerPrivateEndpoints()
}

// registerPrivateEndpoints implements router.
func (r *Router) registerPrivateEndpoints() {
	r.App.Use(r.AuthMiddleware)
	r.App.Get("/secrit", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "This is a secret endpoint",
		})
	})
}

// registerPublicEndpoints implements router.
func (r *Router) registerPublicEndpoints() {
	r.App.Post("/products", r.ProductController.CreateProduct)
	r.App.Get("/products/:id", r.ProductController.GetProductByID)
	r.App.Put("/products/:id", r.ProductController.UpdateProduct)
	r.App.Delete("/products/:id", r.ProductController.DeleteProduct)
}

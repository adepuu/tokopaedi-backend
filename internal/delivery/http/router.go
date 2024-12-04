package http

import "github.com/gofiber/fiber/v2"

type Router struct {
	App               *fiber.App
	ProductController *ProductController
}

type router interface {
	Setup()
	registerPublicEndpoints()
	registerPrivateEndpoints()
}

func NewRouter(app *fiber.App, productController *ProductController) router {
	return &Router{
		App:               app,
		ProductController: productController,
	}
}

// Setup implements router.
func (r *Router) Setup() {
	r.registerPublicEndpoints()
	r.registerPrivateEndpoints()
}

// registerPrivateEndpoints implements router.
func (r *Router) registerPrivateEndpoints() {

}

// registerPublicEndpoints implements router.
func (r *Router) registerPublicEndpoints() {
	r.App.Post("/products", r.ProductController.CreateProduct)
	r.App.Get("/products/:id", r.ProductController.GetProductByID)
	r.App.Put("/products/:id", r.ProductController.UpdateProduct)
	r.App.Delete("/products/:id", r.ProductController.DeleteProduct)
}

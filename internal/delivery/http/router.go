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

func NewRouter(app *fiber.App) router {
	return &Router{
		App: app,
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
}

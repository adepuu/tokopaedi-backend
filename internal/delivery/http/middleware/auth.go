package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func NewAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization", "NOT_FOUND")

		if token == "NOT_FOUND" {
			return fiber.ErrUnauthorized
		}

		ctx.Locals("auth", token)
		return ctx.Next()
	}
}

func getToken(ctx *fiber.Ctx) *string {
	return ctx.Locals("auth").(*string)
}

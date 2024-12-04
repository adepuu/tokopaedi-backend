package config

import (
	"github.com/adepuu/tokopaedi-backend/internal/delivery/http"
	"github.com/adepuu/tokopaedi-backend/internal/repository"
	"github.com/adepuu/tokopaedi-backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB     *gorm.DB
	App    *fiber.App
	Log    *logrus.Logger
	Config *viper.Viper
}

func (cfg *AppConfig) Run() {
	// setup repositories
	productRepository := repository.NewProductRepository(cfg.Log)

	// setup use cases
	productUseCase := usecase.NewProductUsecase(productRepository, cfg.Log, cfg.DB)

	// setup controller
	productController := http.NewProductController(&productUseCase, cfg.Log)

	routeConfig := http.Router{
		App:               cfg.App,
		ProductController: productController,
	}
	routeConfig.Setup()
}

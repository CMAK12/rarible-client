package http_v1

import (
	"inforce-task/internal/dto"

	_ "inforce-task/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"
)

type Service interface {
	GetOwnerships(id string) (map[string]any, error)
	GetTraits(body dto.CollectionTrait) (map[string]any, error)
}

type Handler struct {
	svc    Service
	logger *zap.Logger
}

func NewHandler(svc Service, logger *zap.Logger) *Handler {
	return &Handler{
		svc:    svc,
		logger: logger,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New()

	app.Use(loggerMiddleware(h.logger))
	app.Use(recover.New())
	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	h.initReribleRoutes(app)

	return app
}

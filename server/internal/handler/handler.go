package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"server/internal/config"
	"server/internal/log"
	"server/pkg"

	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "server/docs"
)

type Handler struct {
	db     *sqlx.DB
	logger *zerolog.Logger
}

func NewHandler(db *sqlx.DB, logger *zerolog.Logger) *Handler {
	return &Handler{db: db, logger: logger}
}

func (h *Handler) Router() *fiber.App {
	f := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})

	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		//AllowCredentials: true,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	f.Use(log.RequestLogger(h.logger))

	f.Get("/swagger/*", fiberSwagger.WrapHandler)
	f.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("healthy")
	})
	f.Post("/signup", h.SignUp)
	f.Post("/login", h.Login)

	f.Post("/formula", h.CreateFormula)
	f.Get("/formula/id/:id", h.GetFormulaById)
	f.Get("/formula/user_id/:id", h.GetFormulaByUserId)
	f.Put("/formula", h.UpdateFormula)
	f.Delete("/formula/id/:id", h.DeleteFormula)
	f.Post("/formula/file", h.GetFormulaFromArticle)

	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return pkg.WithJWTAuth(c, config.SigningKey)
	})

	return f
}

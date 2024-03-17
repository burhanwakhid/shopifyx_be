package restserver

import (
	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/handler"
	"github.com/burhanwakhid/shopifyx_backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(fiberServer *fiber.App, handleContainer *handler.HandlerContainer) {
	versionGroup := fiberServer.Group("/v1")

	// /v1/auth/*
	registerAuthRoutes(handleContainer, versionGroup)

	// /v1/bank/*
	registerBankRoutes(handleContainer, versionGroup)

	// /v1/product/*
	registerProductRoutes(handleContainer, versionGroup)

	// /v1/image/*
	registerImageRoutes(handleContainer, versionGroup)
}

func registerAuthRoutes(
	handlerContainer *handler.HandlerContainer,
	group fiber.Router,
) {
	// /v1/auth/*
	auth := group.Group("/auth")
	handlerContainer.AuthHandler.Mount(auth)

}

func registerBankRoutes(
	handlerContainer *handler.HandlerContainer,
	group fiber.Router,
) {
	// /v1/bank/*
	ai := group.Group("/bank")
	ai.Use(middleware.JwtRestAuth())
	handlerContainer.BankHandler.Mount(ai)
}

func registerProductRoutes(
	handlerContainer *handler.HandlerContainer,
	group fiber.Router,
) {
	// /v1/product/*
	ai := group.Group("/product")
	ai.Use(middleware.JwtRestAuth())
	handlerContainer.ProductHandler.Mount(ai)
}

func registerImageRoutes(
	handlerContainer *handler.HandlerContainer,
	group fiber.Router,
) {
	// /v1/image/*
	ai := group.Group("/image")
	ai.Use(middleware.JwtRestAuth())
	handlerContainer.ImageHandler.Mount(ai)
}

package restserver

import (
	"fmt"
	"log"

	"github.com/goccy/go-json"

	"github.com/burhanwakhid/shopifyx_backend/config"
	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Server struct {
	config         *config.Config
	fiberServer    *fiber.App
	fastHTTPServer *fasthttp.Server
}

func NewServer(
	cfg *config.Config,
	handlerContainer *handler.HandlerContainer,
) *Server {

	fiberServer := initializeEchoServer(handlerContainer)

	fastHTTPServer := &fasthttp.Server{
		Handler: fiberServer.Handler(),
	}

	return &Server{
		fiberServer:    fiberServer,
		config:         cfg,
		fastHTTPServer: fastHTTPServer,
	}
}

func (server *Server) Start() error {
	const funcName = "Start"

	addr := fmt.Sprintf(":%d", config.AppPort())

	log.Printf("%s starting REST API server... %s", funcName, addr)

	// err := grace.ServeHTTP(server.fastHTTPServer, httpServer.Addr, 0)

	server.fiberServer.Listen(addr)

	log.Printf("%s REST API server stopped", funcName)

	return nil

}

func initializeEchoServer(
	handlerContainer *handler.HandlerContainer,
) *fiber.App {

	fiberServer := fiber.New(
		fiber.Config{
			Prefork:       true,
			StrictRouting: true,
			JSONEncoder:   json.Marshal,
			JSONDecoder:   json.Unmarshal,
		},
	)

	// Routes.
	// Put this in the last line.
	RegisterRoutes(fiberServer, handlerContainer)

	return fiberServer
}

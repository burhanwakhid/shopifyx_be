//go:build wireinject
// +build wireinject

package restserver

import (
	"github.com/burhanwakhid/shopifyx_backend/config"
	"github.com/google/wire"
)

func InitServer(
	cfg *config.Config,
) *Server {
	wire.Build(allSet)
	return &Server{}
}

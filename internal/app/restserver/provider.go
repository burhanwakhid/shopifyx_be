package restserver

import (
	"github.com/burhanwakhid/shopifyx_backend/internal/app/provider"
	"github.com/google/wire"
)

var (
	// middlewareSet = wire.NewSet(
	// )

	allSet = wire.NewSet(
		provider.ConfigSet,
		provider.PkgSet,
		provider.RepoSet,
		provider.ServiceSet,
		provider.RestDeliverySet,
		NewServer,
	)
)

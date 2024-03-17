package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type LocalOption struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

type Local struct {
	*cache.Cache
}

func NewLocal(
	o LocalOption,
) *Local {
	c := cache.New(
		o.DefaultExpiration,
		o.CleanupInterval,
	)
	return &Local{
		Cache: c,
	}
}

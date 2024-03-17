package main

import (
	"math/rand"
	"time"

	"github.com/burhanwakhid/shopifyx_backend/cmd"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	cmd.Initialize()
}

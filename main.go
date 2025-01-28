package main

import (
	"github.com/marcuschui2022/pokedex/internal/api"
	"time"
)

func main() {
	client := api.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		apiClient:     client,
		caughtPokemon: make(map[string]api.Pokemon),
	}
	startRepl(cfg)
}

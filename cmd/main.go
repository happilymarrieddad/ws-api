package main

import (
	"github.com/happilymarrieddad/ws-api/internal/api"
	"github.com/happilymarrieddad/ws-api/internal/config"
	"github.com/happilymarrieddad/ws-api/internal/wsclient"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	client, err := wsclient.NewWSClient(cfg, nil)
	if err != nil {
		panic(err)
	}

	api.Start(cfg, client, 8000)
}

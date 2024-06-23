package config

import (
	"github.com/hashicorp/consul/api"
	"log"
)

func InitConsulDB() *api.Client {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}
	return client
}

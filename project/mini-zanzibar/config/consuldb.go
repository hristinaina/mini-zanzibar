package config

import (
	"github.com/hashicorp/consul/api"
	"log"
)

func InitConsulDB() *api.Client {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}
	return client
}

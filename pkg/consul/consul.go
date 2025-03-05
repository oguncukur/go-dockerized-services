package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

var consulClient *api.Client

func init() {
	client, err := api.NewClient(&api.Config{
		Address: "localhost:8500",
	})
	if err != nil {
		log.Fatalf("Error creating Consul client: %v", err)
	}
	consulClient = client
}

func TestConsulKV() {

	kv := consulClient.KV()
	pair := &api.KVPair{
		Key:   "my-key",
		Value: []byte("my-value"),
	}

	_, err := kv.Put(pair, nil)
	if err != nil {
		log.Fatalf("Consul KV PUT failed: %v", err)
	}

	pair, _, err = kv.Get("my-key", nil)
	if err != nil {
		log.Fatalf("Consul KV GET failed: %v", err)
	}

	if string(pair.Value) != "my-value" {
		log.Fatalf("Expected 'my-value', got %s", string(pair.Value))
	}

	fmt.Println("Consul KV test passed")
}

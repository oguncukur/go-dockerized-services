package main

import (
	"go-dockerized-services/pkg/consul"
	"go-dockerized-services/pkg/kafka"
	"go-dockerized-services/pkg/redis"
)

func main() {
	consul.TestConsulKV()
	redis.SetAndGet()
	kafka.Producer()
	kafka.Consumer()
}

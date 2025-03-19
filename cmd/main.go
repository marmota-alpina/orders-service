package main

import (
	"github.com/marmota-alpina/orders-service/cmd/graphql"
	"github.com/marmota-alpina/orders-service/cmd/grpc"
	"github.com/marmota-alpina/orders-service/cmd/rest"
	"log"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("🚀 Iniciando API gRPC")
		grpc.StartGrpcServer()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("🚀 Iniciando o GraphQL Server")
		graphql.StartGraphQLServer()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("🚀 Iniciando o API Rest")
		rest.StartAPI()
	}()

	wg.Wait()
}

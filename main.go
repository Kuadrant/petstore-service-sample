package main

import (
	"log"
	"net/http"

	"github.com/kuadrant/petstore-service-sample/pkg/api"
	"github.com/kuadrant/petstore-service-sample/pkg/openapi"
	"github.com/kuadrant/petstore-service-sample/pkg/version"
)

func main() {
	log.Printf("Started PetStore API (Version: %s)\n", version.VERSION)

	go func() {
		log.Fatal(http.ListenAndServe("0.0.0.0:8080", api.GetRouter()))
	}()

	log.Fatal(http.ListenAndServe("0.0.0.0:9090", openapi.GetRouter()))
}

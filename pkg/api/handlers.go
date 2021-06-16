package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/kuadrant/petstore-service-sample/pkg/version"
)

func handleGetPets() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"name":"micky"}, {"name":"minnie"}]`)
	}
}

func constructHealthz() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		out := fmt.Sprintf(`{"version":"%s"}`, version.VERSION)
		fmt.Fprintf(w, out)
	}
}

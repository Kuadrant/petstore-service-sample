package openapi

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed petstore.yaml
var petstoreOpenAPI string

func handleOpenAPI() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/x-yaml")
		fmt.Fprintf(w, petstoreOpenAPI)
	}
}

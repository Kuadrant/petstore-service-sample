package api

import (
	"github.com/julienschmidt/httprouter"
)

func GetRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/healthz", constructHealthz())

	// PET
	router.GET("/pets", handleGetPets())

	return router
}

package openapi

import (
	"github.com/julienschmidt/httprouter"
)

func GetRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/openapi", handleOpenAPI())
	return router
}

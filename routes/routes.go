package routes

import (
	"Tiomon/Pendencias/handler"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", handler.Home)
	r.NoRoute(handler.RoutesNotFound)

	r.Run()
}

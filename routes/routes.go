package routes

import (
	"Tiomon/Pendencias/handler"

	"github.com/gin-gonic/gin"
)

func InicializerRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/pendencias", handler.CreatePendingHandler)
		v1.GET("/pendencias", handler.ShowPedingHandler)
		v1.PUT("/pendencias", handler.UpdatePendingHandler)
		v1.DELETE("/pendencias", handler.DeletePendingHandler)
		v1.GET("/pendencias", handler.ListPendingHandler)

	}
}

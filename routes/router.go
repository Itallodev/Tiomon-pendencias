package routes

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	router.Run(":8080")
}
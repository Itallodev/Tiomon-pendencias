package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePendingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET pending",
	})
}

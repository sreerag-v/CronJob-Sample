package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func SendMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "StatusOk"})
}

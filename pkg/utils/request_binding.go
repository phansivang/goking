package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func BindQuery(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindQuery(obj); err != nil {
		log.Printf("Query binding error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameters"})
		return false
	}
	return true
}

func BindJson(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		log.Printf("Json binding error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Payload"})
		return false
	}
	return true
}

package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, SuccessResponses{Message: "SUCCESS", StatusCode: http.StatusOK})
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Pre-allocate the response structure to avoid repeated allocations
		var response struct {
			Data    interface{} `json:"data"`
			Message string      `json:"message"`
		}

		if c.Writer.Status() == http.StatusOK {
			if responseData, exists := c.Get("Data"); exists {
				response.Data = responseData
			} else {
				response.Data = []interface{}{}
			}
			response.Message = "SUCCESS"
		} else {
			if errorMessage, exists := c.Get("errorMessage"); exists {
				response.Message = errorMessage.(string)
			} else {
				response.Message = "An error occurred"
			}
			response.Data = []interface{}{}
		}

		c.JSON(c.Writer.Status(), response)
	}
}

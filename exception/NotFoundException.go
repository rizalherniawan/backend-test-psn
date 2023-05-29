package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizalherniawan/backend-test-psn/dto/response"
)

func NotFoundException(c *gin.Context, message string) {
	res := response.Error{
		Success: false,
		Message: message,
	}

	c.JSON(http.StatusNotFound, res)
}

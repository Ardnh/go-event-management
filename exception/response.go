package exception

import (
	"go/ems/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, status int, data interface{}, err error) {
	webResponse := domain.WebResponse{
		Code:   status,
		Status: http.StatusText(status),
		Data:   data,
	}

	if err != nil {
		c.JSON(status, &webResponse)
	} else {
		c.JSON(status, &webResponse)
	}
}

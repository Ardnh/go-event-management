package exception

import (
	"go/ems/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, status int, data interface{}, err error) {

	if err != nil {
		webResponse := domain.WebResponse{
			Code:   status,
			Status: http.StatusText(status),
			Data:   err.Error(),
		}
		c.JSON(status, &webResponse)
	} else {
		webResponse := domain.WebResponse{
			Code:   status,
			Status: http.StatusText(status),
			Data:   data,
		}
		c.JSON(status, &webResponse)
	}
}

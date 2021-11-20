package v1

import (
	"github.com/1makarov/go-csv-crud-example/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func newErrorResponse(c *gin.Context, statusCode int, err error) {
	logrus.Error(err)
	c.AbortWithStatusJSON(statusCode, types.ErrorResponse{Error: err.Error()})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorReponse struct {
	Message string `json:"message"`
}

type successReponse struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorReponse{message})
}

func newSuccessResponse(c *gin.Context, message string, result interface{}) {
	c.JSON(http.StatusOK, successReponse{"successfully " + message, result})
}

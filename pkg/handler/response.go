package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{message}) // принимает http-статус код и тело ответа,
	// в качестве ответа функция принимает интерфейс,
	// поэтому мы можем передать как структуру,
	// так и мапу со строкой в качестве ключа и интерфейса в качестве значения
}

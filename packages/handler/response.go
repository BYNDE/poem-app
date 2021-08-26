package handler

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newResponse(c *gin.Context, StatusCode int, obj interface{}) {
	req, _ := url.QueryUnescape(c.Request.RequestURI)
	logrus.Infoln("Method: '" + c.Request.Method + "' RequestURI: '" + req + "' Status: '" + strconv.Itoa(StatusCode) + " " + http.StatusText(StatusCode) + "'")
	c.JSON(StatusCode, obj)
}

func newErrorResponse(c *gin.Context, StatusCode int, message string) {
	req, _ := url.QueryUnescape(c.Request.RequestURI)
	logrus.Error("Method: '" + c.Request.Method + "' RequestURI: '" + req + "' Status: '" + strconv.Itoa(StatusCode) + " " + http.StatusText(StatusCode) + "' Message: '" + message + "'")
	c.AbortWithStatusJSON(StatusCode, errorResponse{Message: message})
}
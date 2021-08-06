package handler

import (
	"errors"
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
	ip, _ := getClientIPByHeaders(c.Request)
	req, _ := url.QueryUnescape(c.Request.RequestURI)
	logrus.Infoln("Method: '" + c.Request.Method + "' RequestURI: '" + req + "' Status: '" + strconv.Itoa(StatusCode) + " " + http.StatusText(StatusCode) + "' CLientIp: '" + ip + "'")
	c.JSON(StatusCode, obj)
}

func newErrorResponse(c *gin.Context, StatusCode int, message string) {
	ip, _ := getClientIPByHeaders(c.Request)
	req, _ := url.QueryUnescape(c.Request.RequestURI)
	logrus.Error("Method: '" + c.Request.Method + "' RequestURI: '" + req + "' Status: '" + strconv.Itoa(StatusCode) + " " + http.StatusText(StatusCode) + "' Message: '" + message + "' CLientIp: '" + ip + "'")
	c.AbortWithStatusJSON(StatusCode, errorResponse{Message: message})
}

func getClientIPByHeaders(req *http.Request) (ip string, err error) {

	// Client could be behid a Proxy, so Try Request Headers (X-Forwarder)
	ipSlice := []string{}

	ipSlice = append(ipSlice, req.Header.Get("X-Forwarded-For"))
	ipSlice = append(ipSlice, req.Header.Get("x-forwarded-for"))
	ipSlice = append(ipSlice, req.Header.Get("X-FORWARDED-FOR"))

	for _, v := range ipSlice {
		if v != "" {
			return v, nil
		}
	}
	err = errors.New("error: Could not find clients IP address from the Request Headers")
	return "Unknown", err

}

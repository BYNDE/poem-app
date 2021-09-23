package handler

import (
	"github.com/dvd-denis/IT-Platform/packages/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) Handler {
	return Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api") // ! h.userIdentity - индификация пользователя
	{
		platforms := api.Group("/platforms")
		{
			platforms.POST("/", h.addPlatform)
			platforms.GET(":id", h.getPlatformById)
			platforms.GET("limit/:limit", h.GetAllPlatformsLimit)
			platforms.PUT(":id", h.updatePlatform)
			platforms.DELETE(":id", h.deletePlatform)
			platforms.GET("title/:title", h.getPlatformByTitle)
		}
		authors := api.Group("/authors")
		{
			authors.POST("/", h.addAuthor)
			authors.PUT(":id", h.updateAuthor)
			authors.GET(":id", h.getAuthorById)
			authors.GET(":id/platforms", h.getPlatformsById)
			authors.GET("limit/:limit", h.GetAllAuthorsLimit)
			authors.GET("name/:name", h.getAuthorByTitle)
			authors.DELETE(":id", h.deleteAuthor)
		}
	}

	return router
}

package handler

import (
	"github.com/dvd-denis/poem-app/packages/service"
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
		poems := api.Group("/poems")
		{
			poems.POST("/", h.addPoem)
			poems.GET(":id", h.getPoemById)
			poems.GET("limit/:limit", h.GetAllPoemsLimit)
			poems.PUT(":id", h.updatePoem)
			poems.DELETE(":id", h.deletePoem)
			poems.GET("title/:title", h.getPoemByTitle)
		}
		authors := api.Group("/authors")
		{
			authors.POST("/", h.addAuthor)
			authors.PUT(":id", h.updateAuthor)
			authors.GET(":id", h.getAuthorById)
			authors.GET(":id/poems", h.getPoemsById)
			authors.GET("limit/:limit", h.GetAllAuthorsLimit)
			authors.GET("name/:name", h.getAuthorByTitle)
			authors.DELETE(":id", h.deleteAuthor)
		}
	}

	return router
}

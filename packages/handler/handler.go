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

	api := router.Group("/api", h.userIdentity)
	{
		poems := api.Group("/poems")
		{
			poems.POST("/", h.addPoem)
			poems.GET("/", h.getAllPoems) // ! удалить при реальной базе данных
			poems.GET(":id", h.getPoemById)
			poems.PUT(":id", h.updatePoem)
			poems.DELETE(":id", h.deletePoem)
			poems.GET("title/:title", h.getPoemByTitle)
		}
		authors := api.Group("/authors")
		{
			authors.POST("/", h.addAuthor)
			authors.GET("/", h.getAllAuthors) // ! удалить при реальной базе данных
			authors.PUT(":id", h.updateAuthor)
			authors.GET(":id", h.getAuthorById)
			authors.DELETE(":id", h.deleteAuthor)
		}
	}

	return router
}

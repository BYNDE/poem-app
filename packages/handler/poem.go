package handler

import (
	"net/http"

	"github.com/dvd-denis/poem-app"
	"github.com/gin-gonic/gin"
)

type poemInInput struct {
	Title    string `json:"title" binding:"required"`
	Text     string `json:"text" binding:"required"`
	AuthorId int    `json:"authorId"`
}

func (h *Handler) addPoem(c *gin.Context) {
	var input poemInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	poem := poem.Poems{
		Title: input.Title,
		Text:  input.Text,
	}

	id, err := h.services.Poem.Create(input.AuthorId, poem)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getPoemById(c *gin.Context) {

}

func (h *Handler) getPoemByTitle(c *gin.Context) {

}

func (h *Handler) updatePoem(c *gin.Context) {

}

func (h *Handler) deletePoem(c *gin.Context) {

}

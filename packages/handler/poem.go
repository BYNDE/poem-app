package handler

import (
	"net/http"

	"github.com/dvd-denis/poem-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addPoem(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	var input poem.Poems
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.AuthorId == 0 {
		input.AuthorId = 1
	}

	id, err := h.services.Poem.Create(input)
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

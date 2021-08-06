package handler

import (
	"net/http"
	"strconv"

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

	newResponse(c, http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getPoemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	poem, err := h.services.Poem.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, poem)
}

// type getAllPoemsResponse struct {
// 	Data []poem.Poems `json:"data"`
// }

func (h *Handler) getPoemByTitle(c *gin.Context) {
	title := c.Param("title")

	poems, err := h.services.Poem.GetByTitle(title)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, poems)
}

func (h *Handler) updatePoem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input poem.UpdatePoemInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Poem.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deletePoem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Poem.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, statusResponse{"ok"})
}

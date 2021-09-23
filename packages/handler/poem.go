package handler

import (
	"net/http"
	"strconv"

	platform "github.com/dvd-denis/IT-Platform"
	"github.com/gin-gonic/gin"
)

type platformInInput struct {
	Title    string `json:"title" binding:"required"`
	Text     string `json:"text" binding:"required"`
	AuthorId int    `json:"authorId"`
}

func (h *Handler) addPlatform(c *gin.Context) {
	var input platformInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	platform := platform.Platforms{
		Title: input.Title,
		Text:  input.Text,
	}

	id, err := h.services.Platform.Create(input.AuthorId, platform)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllPlatformsLimit(c *gin.Context) {
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	platforms, err := h.services.Platform.GetAllLimit(limit)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, http.StatusOK, platforms)
}

func (h *Handler) getPlatformById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	platform, err := h.services.Platform.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, platform)
}

// type getAllPlatformsResponse struct {
// 	Data []platform.Platforms `json:"data"`
// }

func (h *Handler) getPlatformByTitle(c *gin.Context) {
	title := c.Param("title")

	platforms, err := h.services.Platform.GetByTitle(title)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, platforms)
}

func (h *Handler) updatePlatform(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input platform.UpdatePlatformInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Platform.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deletePlatform(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Platform.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, statusResponse{"ok"})
}

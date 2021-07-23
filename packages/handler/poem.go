package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addPoem(c *gin.Context) {
	// _, ok := c.Get(userCtx)
	// if !ok {
	// 	newErrorResponse(c, http.StatusInternalServerError, "user id not found")
	// 	return
	// }
	// var input poem.Poems
	// if err := c.BindJSON(&input); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	id, _ := c.Get(userCtx)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	// ! работа с сервисом
}

func (h *Handler) getAllPoems(c *gin.Context) {

}

func (h *Handler) getPoemById(c *gin.Context) {

}

func (h *Handler) getPoemByTitle(c *gin.Context) {

}

func (h *Handler) updatePoem(c *gin.Context) {

}

func (h *Handler) deletePoem(c *gin.Context) {

}

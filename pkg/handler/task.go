package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test"
)

func (h *Handler) getAllTasks(c *gin.Context) {
	list, err := h.services.GetTaskList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) addTask(c *gin.Context) {
	var input test.AddTask

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.AddTask(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type responseReminder struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func (h Handler) GetReminder(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "invalid id")
		return
	}
	reminder, err := h.repository.Get(id)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
	}
	context.JSON(http.StatusOK, responseReminder{Id: reminder.Id, Task: reminder.Task, Done: reminder.Done})
}

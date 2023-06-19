package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type newReminder struct {
	Task string `json:"task"`
}

type createdReminder struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
}

func (h Handler) CreateReminder(context *gin.Context) {
	reminder := newReminder{}
	err := context.BindJSON(&reminder)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := h.repository.Create(reminder.Task)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, createdReminder{id, reminder.Task})
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type changeReminder struct {
	Id int `json:"id"`
}

type reminderResponse struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func (h Handler) UpdateReminder(context *gin.Context) {
	reminder := changeReminder{}
	err := context.BindJSON(&reminder)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	updatedReminder, err := h.repository.Update(reminder.Id)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, reminderResponse{updatedReminder.Id, updatedReminder.Task, updatedReminder.Done})
}

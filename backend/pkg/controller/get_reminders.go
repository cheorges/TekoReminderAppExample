package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type remindersResponse struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func (h Handler) GetReminders(context *gin.Context) {
	reminders := h.repository.GetAll()
	remindersResponses := make([]remindersResponse, len(reminders))
	for index, reminder := range reminders {
		remindersResponses[index] = remindersResponse{Id: reminder.Id, Task: reminder.Task, Done: reminder.Done}
	}
	context.JSON(http.StatusOK, remindersResponses)
}

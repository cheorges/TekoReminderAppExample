package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) DeleteReminder(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	deleted := h.repository.Delete(id)

	if deleted {
		context.Status(http.StatusOK)
	} else {
		context.JSON(http.StatusBadRequest, "could not delete reminder")
	}
}

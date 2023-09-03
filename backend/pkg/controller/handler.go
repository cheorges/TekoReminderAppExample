package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"reminders/pkg/repository"
)

type Handler struct {
	engine     *gin.Engine
	repository *repository.PostgresqlStorageImpl
}

func New() *Handler {
	h := &Handler{
		engine:     gin.Default(),
		repository: repository.NewPostgresStorageImpl(),
	}

	configureCors(h.engine)

	router := h.engine.Group("/api/reminder/")
	router.GET("/:id", h.GetReminder)
	router.DELETE("/:id", h.DeleteReminder)
	router.GET("/all", h.GetReminders)
	router.POST("/", h.CreateReminder)
	router.PUT("/", h.UpdateReminder)
	return h
}

func (h Handler) Run() {
	err := h.engine.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func configureCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	engine.Use(cors.New(config))
}

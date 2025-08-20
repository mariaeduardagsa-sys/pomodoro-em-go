package handlers

import (
	"net/http"
	"time"

	"github.com/mariaeduardagsa-sys/pomodoro-em-go/pomodoro"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Pomo *pomodoro.Pomodoro
}

func NewHandler(p *pomodoro.Pomodoro) *Handler {
	return &Handler{Pomo: p}
}

func (h *Handler) Start(c *gin.Context) {
	h.Pomo.Start(25 * time.Minute)
	c.JSON(200, gin.H{"status": "Pomodoro iniciado por 25 minutos"})
}

func (h *Handler) Status(c *gin.Context) {
	remaining, running := h.Pomo.Status()
	if !running {
		c.JSON(200, gin.H{"status": "Nenhum Pomodoro em andamento"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		" remaining": int(remaining.Minutes()),
		"running":    running,
	})
}

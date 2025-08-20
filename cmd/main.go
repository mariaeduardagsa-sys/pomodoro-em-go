package main

import (
	"github.com/mariaeduardagsa-sys/pomodoro-em-go/handlers"
	"github.com/mariaeduardagsa-sys/pomodoro-em-go/pomodoro"

	"github.com/gin-gonic/gin"
)

func main() {

	p := pomodoro.New()
	h := handlers.NewHandler(p)

	r := gin.Default()
	r.POST("/start", h.Start)
	r.GET("/status", h.Status)
	r.Run()
}

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learnda1ly/homeschool_tracker/school"
)

func main() {
	data := struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}{
		Status: "ok",
		Data:   "server running",
	}

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
	})

	r.GET("/learners", func(c *gin.Context) {
		c.JSON(http.StatusOK, school.GetLearnerList())
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("./public" + c.Request.URL.Path)
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

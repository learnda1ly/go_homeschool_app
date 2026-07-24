package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func main() {
	r := gin.Default()

	// 1. Register API routes FIRST
	api := r.Group("/api")
	{
		api.GET("/ping", pingHandler)
		// ... all your JSON endpoints
	}

	// 2. Catch-all for static files at root (after APIs)
	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
			return
		}

		// Safe static serving from public/
		c.FileFromFS(c.Request.URL.Path, http.Dir("./public"))
	})

	_ = r.Run(":8080")

	/*
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			jsonBytes, _ := json.Marshal(data)
			_, _ = w.Write(jsonBytes)
		})

		http.Handle("/", http.FileServer(http.Dir("./public")))

		server := &http.Server{
			Addr:         ":8080",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	*/
}

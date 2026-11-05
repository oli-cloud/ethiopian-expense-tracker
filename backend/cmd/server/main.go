package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oli/birrwise-backend/internal/config"
	"github.com/oli/birrwise-backend/internal/database"
	"github.com/oli/birrwise-backend/internal/handlers"
)

func main() {
	cfg := config.Load()
	gin.SetMode(cfg.GinMode)
	db, err := database.Open(cfg.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), corsMiddleware(cfg.CORSOrigins))
	handlers.Register(r, db)
	log.Printf("BirrWise API listening on http://localhost:%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

func corsMiddleware(origins string) gin.HandlerFunc {
	allowed := strings.Split(origins, ",")
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		for _, item := range allowed {
			if strings.TrimSpace(item) == origin {
				c.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

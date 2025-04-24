package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"api-golang/database"
	"strings"
)

func init() {
	databaseUrl := os.Getenv("DATABASE_URL")

	// If not set via env (dev), fallback to secret file (prod)
	if databaseUrl == "" {
		content, err := os.ReadFile("/run/secrets/DATABASE_URL")
		if err != nil {
			log.Fatal(err)
		}
		databaseUrl = strings.TrimSpace(string(content))
	}

	errDB := database.InitDB(databaseUrl)
	if errDB != nil {
		log.Fatalf("⛔ Unable to connect to database: %v\n", errDB)
	} else {
		log.Println("DATABASE CONNECTED 🥇")
	}

}

func main() {

	r := gin.Default()
	var tm time.Time

	r.GET("/", func(c *gin.Context) {
		tm = database.GetTime(c)
		c.JSON(200, gin.H{
			"api": "golang",
			"now": tm,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		tm = database.GetTime(c)
		c.JSON(200, "pong")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (or "PORT" env var)
}

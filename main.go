package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thirst154/url-shorter/models"
)

func main() {
	println("Starting server...")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	if secret := os.Getenv("POSTGRES_CONNECTION"); secret == "" {
		log.Fatal("POSTGRES_CONNECTION must exist")
	}

	println("Connecting to DB...")
	models.ConnectDB()

	r := gin.Default()

	// load html files in the current directory
	r.LoadHTMLGlob("*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	// Load favicon
	r.StaticFile("/favicon.ico", "./favicon.svg")

	r.Run(":3000")
}

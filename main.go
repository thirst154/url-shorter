package main

import (
	"log"
	neturl "net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thirst154/url-shorter/models"
	"github.com/thirst154/url-shorter/utils"
)

func main() {
	log.Println("Starting server...")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	if os.Getenv("POSTGRES_CONNECTION") == "" {
		log.Fatal("POSTGRES_CONNECTION must be set")
	}

	log.Println("Connecting to DB...")
	models.ConnectDB()

	r := gin.Default()
	r.LoadHTMLGlob("*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.StaticFile("/favicon.ico", "./favicon.svg")

	r.POST("/shorten", createShortURL)
	r.GET("/:code", redirectURL)

	log.Fatal(r.Run(":3000"))
}

type ShortURLInput struct {
	Code        string `json:"code"         binding:"omitempty,alphanum,max=16"`
	OriginalURL string `json:"original_url" binding:"required,url"`
}

func createShortURL(c *gin.Context) {
	var input ShortURLInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	parsed, err := neturl.Parse(input.OriginalURL)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		c.JSON(400, gin.H{"error": "URL must use http or https"})
		return
	}

	code := input.Code
	if code == "" {
		for {
			code, err = utils.GenerateCode(8)
			if err != nil {
				c.JSON(500, gin.H{"error": "failed to generate code"})
				return
			}
			if models.IsCodeUnique(code) {
				break
			}
		}
	} else if !models.IsCodeUnique(code) {
		c.JSON(409, gin.H{"error": "code already in use"})
		return
	}

	expiresAt := time.Now().Add(time.Hour * 24 * 30 * 6)
	record, err := models.CreateURL(code, input.OriginalURL, &expiresAt)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create short URL"})
		return
	}

	c.JSON(201, gin.H{"code": record.Code})
}

func redirectURL(c *gin.Context) {
	code := c.Param("code")
	record, err := models.GetURL(code)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	if record.ExpiresAt != nil && time.Now().After(*record.ExpiresAt) {
		c.JSON(410, gin.H{"error": "URL has expired"})
		return
	}

	models.IncrementClicks(record.ID)
	c.Redirect(302, record.OriginalURL)
}

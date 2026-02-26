package main

import "github.com/gin-gonic/gin"

func main() {
	println("Starting server...")
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

package main

import (
	"fmt"
	config "hihand_shortener/core"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func encodeBase62(num int64, chars string) string {
	if num == 0 {
		return string(chars[0])
	}
	result := ""
	for num > 0 {
		result = string(chars[num%62]) + result
		num /= 62
	}
	return result
}

func homeHandler(c *gin.Context) {
	// Render the home page without passing data (no need for tmpl.Execute)
	c.HTML(200, "index.html", nil)
}

func shortenHandler(c *gin.Context) {
	url := c.DefaultPostForm("url", "")
	if url == "" {
		c.JSON(400, gin.H{"error": "URL is required"})
		return
	}
	id, _ := rdb.Incr(ctx, "url_counter").Result()
	code := encodeBase62(id, config.AppConfig.BASE62_CHARS)
	rdb.Set(ctx, code, url, 0)
	shortURL := fmt.Sprintf("%s/%s", config.AppConfig.BASE_SHORT_URL, code)
	c.HTML(200, "index.html", gin.H{"ShortURL": shortURL})
}

func redirectHandler(c *gin.Context) {
	code := c.Param("code")
	url, err := rdb.Get(ctx, code).Result()
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	c.Redirect(302, url)
}

func main() {
	config.InitConfig(".")

	// Create Gin router
	r := gin.Default()

	// Serve static files (e.g., favicon, images, CSS, JS)
	r.Static("/static", "./public")

	// Load HTML templates using Gin
	r.LoadHTMLFiles(filepath.Join("templates", "index.html"))

	// Connect Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.AppConfig.REDIS_HOST, config.AppConfig.REDIS_PORT),
	})

	// Set up routes
	r.GET("/", homeHandler)
	r.POST("/shorten", shortenHandler)
	r.GET("/r/:code", redirectHandler)

	// Start the server
	log.Printf("Server started at :%s\n", config.AppConfig.BASE_URL)
	r.Run(":" + config.AppConfig.APP_PORT)
}

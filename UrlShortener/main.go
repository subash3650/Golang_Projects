package main

import (
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Url struct {
	OriginalUrl  string `json:"original_url"`
	ShortenedUrl string `json:"shortened_url,omitempty"`
}

var urlData = make(map[string]string)
var mutex sync.RWMutex

func main() {
	r := gin.Default()
	r.GET("/", HandleRoot)
	r.POST("/shorten", HandleShortenURL)
	r.GET("/:code", HandleCode)
	r.Run(":8080")
}

func HandleCode(c *gin.Context) {
	code := c.Param("code")

	mutex.RLock()
	originalURL, ok := urlData[code]
	mutex.RUnlock()

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func HandleRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the URL Shortener API",
		"status":  "success",
	})
}

func HandleShortenURL(c *gin.Context) {
	var req Url
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format"})
		return
	}
	code := GenerateCode()
	mutex.Lock()
	defer mutex.Unlock()
	req.ShortenedUrl = code
	urlData[code] = req.OriginalUrl
	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + code,
	})
}

func GenerateCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, 10)
	for i := range code {
		code[i] = letters[r.Intn(len(letters))]
	}
	return string(code)
}

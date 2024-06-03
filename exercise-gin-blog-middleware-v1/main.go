package main

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini."},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini."},
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		auth := authHeader[len("Basic "):]
		decodedAuth, err := base64.StdEncoding.DecodeString(auth)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		credentials := string(decodedAuth)
		for _, user := range users {
			if credentials == user.Username+":"+user.Password {
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
	}
}

func getCurrentTime() time.Time {
	// Menghasilkan waktu saat ini
	return time.Now().UTC()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(authMiddleware())

	r.GET("/posts", func(c *gin.Context) {
		idParam := c.Query("id")
		if idParam != "" {
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
				return
			}

			var foundPost *Post
			for _, post := range Posts {
				if post.ID == id {
					foundPost = &post
					break
				}
			}

			if foundPost == nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"post": foundPost})
		} else {
			c.JSON(http.StatusOK, gin.H{"posts": Posts})
		}
	})

	r.POST("/posts", func(c *gin.Context) {
		var newPost Post
		if err := c.ShouldBindJSON(&newPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		newPost.ID = len(Posts) + 1
		newPost.CreatedAt = getCurrentTime()
		newPost.UpdatedAt = getCurrentTime()

		Posts = append(Posts, newPost)

		c.JSON(http.StatusCreated, gin.H{"message": "Postingan berhasil ditambahkan", "post": newPost})
	})

	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}

// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)

type Post struct {
	ID      int
	Title   string
	Content string
}

var posts []Post

func main() {
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("core/templates/*")

	// Define routes
	router.GET("/", showPosts)
	router.GET("/post/:id", showPost)
	router.POST("/post", createPost)

	router.Run(":8080")
}

func showPosts(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"posts": posts,
	})
}

func showPost(c *gin.Context) {
    // Retrieve post ID from URL
    postID := c.Param("id")

    // Convert postID to an integer
    id, err := strconv.Atoi(postID)
    if err != nil {
        // Handle the error, for example, show an error page or message
        c.String(http.StatusBadRequest, "Invalid post ID")
        return
    }

    // Find the post with the given ID
    var post Post
    for _, p := range posts {
        if p.ID == id {
            post = p
            break
        }
    }

    c.HTML(200, "post.tmpl", gin.H{
        "post": post,
    })
}


func createPost(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	// Generate a unique ID for the new post
	newID := len(posts) + 1

	// Create a new post and add it to the posts slice
	newPost := Post{
		ID:      newID,
		Title:   title,
		Content: content,
	}
	posts = append(posts, newPost)

	// Redirect to the homepage to show the new post
	c.Redirect(302, "/")
}

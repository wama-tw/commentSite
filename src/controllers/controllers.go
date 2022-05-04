package controllers

import (
	"net/http"
	"strings"
	"time"

	"OSProject1/src/models"

	"github.com/gin-gonic/gin"
)

var posts = make(map[string]models.Post)

func slug(title string) (slug string) {
	slug = strings.ToLower(title)
	slug = strings.Replace(slug, " ", "-", -1)

	return slug
}

func GetCookies(c *gin.Context) (name string) {
	data, err := c.Cookie("userName")
	if err == nil && data != "" {
		name = data
	} else {
		name = "Anonymous"
	}
	return name
}

func GetAllPosts(c *gin.Context) {
	if _, ok := posts["hello-world"]; !ok {
		posts["hello-world"] = models.Post{
			Title:     "Hello World",
			Author:    "wama",
			Content:   "文章排序是用 Parallel Merge Sort",
			Slug:      "hello-world",
			CreatedAt: time.Now(),
		}
		posts["ping-pong"] = models.Post{
			Title:     "Ping",
			Author:    "wama",
			Content:   "pong pong pong pong pong pong",
			Slug:      "ping-pong",
			CreatedAt: time.Now(),
		}
	}

	var sortedPosts []models.Post
	for _, v := range posts {
		sortedPosts = append(sortedPosts, v)
	}
	mergeSortChan := make(chan []models.Post)
	go mergeSort(sortedPosts, mergeSortChan)
	sortedPosts = <-mergeSortChan

	c.HTML(http.StatusOK, "index.html", gin.H{
		"posts": sortedPosts,
		"name":  GetCookies(c),
	})
}

func GetCreatePost(c *gin.Context) {
	c.HTML(http.StatusOK, "newPost.html", gin.H{
		"name": GetCookies(c),
	})
}

func CreatePost(c *gin.Context) {
	newPostTitle := c.PostForm("title")
	newPostSlug := slug(newPostTitle)
	posts[newPostSlug] = models.Post{
		Title:     newPostTitle,
		Content:   c.PostForm("content"),
		Author:    GetCookies(c),
		Slug:      newPostSlug,
		CreatedAt: time.Now(),
	}

	c.Redirect(http.StatusFound, "/posts")
}

func GetNaming(c *gin.Context) {
	c.HTML(http.StatusOK, "naming.html", gin.H{
		"name": GetCookies(c),
	})
}

func Naming(c *gin.Context) {
	c.SetCookie("userName", c.PostForm("name"), 3600, "/", "localhost", false, false)
	c.Redirect(http.StatusFound, "/posts")
}

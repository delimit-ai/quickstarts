package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var (
	tasks  = make(map[int]Task)
	nextID = 1
	mu     sync.Mutex
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/tasks", func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		list := make([]Task, 0, len(tasks))
		for _, t := range tasks {
			list = append(list, t)
		}
		c.JSON(http.StatusOK, list)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var body struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mu.Lock()
		defer mu.Unlock()
		task := Task{ID: nextID, Title: body.Title, Description: body.Description, Done: false}
		tasks[nextID] = task
		nextID++
		c.JSON(http.StatusCreated, task)
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		mu.Lock()
		defer mu.Unlock()
		task, ok := tasks[id]
		if !ok {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, task)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		mu.Lock()
		defer mu.Unlock()
		if _, ok := tasks[id]; !ok {
			c.Status(http.StatusNotFound)
			return
		}
		delete(tasks, id)
		c.Status(http.StatusNoContent)
	})

	r.Run(":8080")
}

package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

var tasks = []Task{}

func createTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = uuid.New().String()
	newTask.CreatedAt = time.Now()
	newTask.Completed = false

	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")
	for _, t := range tasks {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updated Task
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Title = updated.Title
			tasks[i].Description = updated.Description
			tasks[i].Completed = updated.Completed
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func main() {
	r := gin.Default()

	r.POST("/tasks", createTask)
	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTaskByID)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	r.Run(":8080")
}

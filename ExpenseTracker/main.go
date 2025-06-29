package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Expense struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

var (
	expense = make(map[int]Expense)
	mutex   sync.RWMutex
	nextID  = 1
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Expense tracker is running")
	})
	r.POST("/expense", createExpense)
	r.GET("/expense", getExpense)
	r.GET("/expense/:id", getExpenseByID)
	r.PUT("/expense/:id", updateExpense)
	r.DELETE("/expense/:id", deleteExpense)
	r.Run(":8080")
}

func createExpense(c *gin.Context) {
	var newExpense Expense
	if err := c.BindJSON(&newExpense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newExpense.Title == "" || newExpense.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and amount are required"})
		return
	}
	mutex.Lock()
	newExpense.ID = nextID
	newExpense.Date = time.Now()
	expense[nextID] = newExpense
	nextID++
	mutex.Unlock()
	c.JSON(http.StatusCreated, newExpense)
}

func getExpense(c *gin.Context) {
	mutex.RLock()
	defer mutex.RUnlock()
	var list []Expense
	for _, e := range expense {
		list = append(list, e)
	}
	c.JSON(http.StatusOK, list)
}

func getExpenseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	mutex.RLock()
	e, ok := expense[id]
	mutex.RUnlock()
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}
	c.JSON(http.StatusOK, e)
}

func updateExpense(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updated Expense
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if updated.Title == "" || updated.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and amount are required"})
		return
	}
	mutex.Lock()
	e, ok := expense[id]
	if !ok {
		mutex.Unlock()
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}
	e.Title = updated.Title
	e.Amount = updated.Amount
	e.Category = updated.Category
	e.Description = updated.Description
	expense[id] = e
	mutex.Unlock()
	c.JSON(http.StatusOK, e)
}

func deleteExpense(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	mutex.Lock()
	_, ok := expense[id]
	if !ok {
		mutex.Unlock()
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}
	delete(expense, id)
	mutex.Unlock()
	c.Status(http.StatusNoContent)
}

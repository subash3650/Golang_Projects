package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

var (
	Client     *mongo.Client
	Collection *mongo.Collection
)

// var tasks = []Task{}

func createTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = uuid.New().String()
	newTask.CreatedAt = time.Now()
	newTask.Completed = false

	_, err := Collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func getTasks(c *gin.Context) {
	cursor, err := Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	defer cursor.Close(context.TODO())

	var tasks []Task
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)

}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task Task
	err := Collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve the user"})
		return
	}

	c.JSON(http.StatusOK, task)

}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updated Task

	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"title":       updated.Title,
			"description": updated.Description,
			"completed":   updated.Completed,
		},
	}

	res, err := Collection.UpdateOne(context.TODO(), bson.M{"id": id}, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update the task"})
		return
	}

	if res.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "User updated successfully"})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	res, err := Collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	con := godotenv.Load()
	if con != nil {
		log.Fatal("error loading .env file")
	}
	uri := os.Getenv("MONGO_URI")

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	Collection = Client.Database("taskTracker").Collection("tasks")

	r := gin.Default()

	r.POST("/tasks", createTask)
	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTaskByID)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	r.Run(":8080")
}

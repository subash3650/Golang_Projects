// main.go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID          string    `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Completed   bool      `json:"completed" bson:"completed"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}

var (
	Client     *mongo.Client
	Collection *mongo.Collection
)

func createTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = uuid.New().String()
	newTask.CreatedAt = time.Now()
	if _, err := Collection.InsertOne(context.TODO(), newTask); err != nil {
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
	if err := cursor.All(context.TODO(), &tasks); err != nil {
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
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve the task"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	var updated Task
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	res, err := Collection.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.M{"$set": bson.M{
			"title":       updated.Title,
			"description": updated.Description,
			"completed":   updated.Completed,
		}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update the task"})
		return
	}
	if res.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
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
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file (continuing with env vars)")
	}
	uri := os.Getenv("MONGO_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	Collection = Client.Database("taskTracker").Collection("tasks")

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/tasks", createTask)
	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTaskByID)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)
	r.Run(":8080")
}

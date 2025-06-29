// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"sync"
// )

// type User struct {
// 	Name string `json:"name"`
// }

// var userData = make(map[int]string)
// var mutex sync.RWMutex

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", welcome)
// 	mux.HandleFunc("/users", createUser)
// 	mux.HandleFunc("/user/", userHandler) // all /user/{id} requests go here

// 	fmt.Println("Server is listening on port 8080")
// 	err := http.ListenAndServe(":8080", mux)
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}
// }

// func welcome(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Task Tracker API is running! Use POST /users to create a user.")
// }

// func createUser(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	if user.Name == "" {
// 		http.Error(w, "Name is required", http.StatusBadRequest)
// 		return
// 	}
// 	mutex.Lock()
// 	userData[len(userData)+1] = user.Name
// 	mutex.Unlock()
// 	w.WriteHeader(http.StatusNoContent)
// 	fmt.Println("User created Successfully")
// }

// func userHandler(w http.ResponseWriter, r *http.Request) {
// 	pathParts := strings.Split(r.URL.Path, "/")
// 	if len(pathParts) != 3 {
// 		http.Error(w, "Invalid URL", http.StatusBadRequest)
// 		return
// 	}
// 	idStr := pathParts[2]
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	switch r.Method {
// 	case http.MethodGet:
// 		getUser(w, id)
// 	case http.MethodDelete:
// 		deleteUser(w, id)
// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func getUser(w http.ResponseWriter, id int) {
// 	mutex.RLock()
// 	user, ok := userData[id]
// 	mutex.RUnlock()
// 	if !ok {
// 		http.Error(w, "user not found", http.StatusNotFound)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	j, err := json.Marshal(user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(j)
// }

// func deleteUser(w http.ResponseWriter, id int) {
// 	mutex.Lock()
// 	defer mutex.Unlock()
// 	_, ok := userData[id]
// 	if !ok {
// 		http.Error(w, "user not found", http.StatusNotFound)
// 		return
// 	}
// 	delete(userData, id)
// 	w.WriteHeader(http.StatusNoContent)
// }

package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
}

var userData = make(map[int]string)
var mutex sync.RWMutex

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Task Tracker API is running! Use POST /users to create a user.")
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if user.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
			return
		}
		mutex.Lock()
		id := len(userData) + 1
		userData[id] = user.Name
		mutex.Unlock()
		c.JSON(http.StatusCreated, gin.H{"id": id, "name": user.Name})
	})

	r.GET("/user/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		mutex.RLock()
		user, ok := userData[id]
		mutex.RUnlock()
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": user})
	})

	r.DELETE("/user/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		_, ok := userData[id]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		delete(userData, id)
		c.Status(http.StatusNoContent)
	})
	r.Run(":8080")
}
 
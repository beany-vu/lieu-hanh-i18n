// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

var mongoSession *mgo.Session

func main() {
	// Connect to MongoDB
	var err error
	mongoSession, err = mgo.Dial("mongodb://mongodb:27018")
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer mongoSession.Close()

	r := gin.Default()

	// Use CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/list-context", listContext)

	// Start the server
	r.Run(":8080")
}

// List the items inside the MongoDB collection "context"
func listContext(ctx *gin.Context) {
	// Retrieve data from MongoDB sampledb collection
	collection := mongoSession.DB("lieu-hanh-i18n").C("context")

	// Perform MongoDB queries here (e.g., collection.Find(...))
	// Example:
	var results []interface{} // Change this to the type of your data
	err := collection.Find(nil).All(&results)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": results})
}

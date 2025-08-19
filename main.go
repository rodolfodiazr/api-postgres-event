package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rodolfodiazr/api-postgres-event/databases"
	"github.com/rodolfodiazr/api-postgres-event/handlers"
	"github.com/rodolfodiazr/api-postgres-event/middlewares"
)

func main() {
	db, err := databases.Connect()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(middlewares.Database(db))

	eventHandler := handlers.EventHandler{}
	router.GET("/events", eventHandler.List)
	router.GET("/events/:id", eventHandler.Get)
	router.POST("/events", eventHandler.Create)

	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}

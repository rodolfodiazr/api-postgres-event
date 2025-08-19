package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/rodolfodiazr/api-postgres-event/models"
	"github.com/rodolfodiazr/api-postgres-event/repositories"
	"github.com/rodolfodiazr/api-postgres-event/services"

	"github.com/gin-gonic/gin"
)

type EventHandler struct{}

func (e EventHandler) List(c *gin.Context) {
	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not found"})
		return
	}

	service := services.NewEventService(repositories.NewEventRepository(db.(*sql.DB)))
	events, err := service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

func (e EventHandler) Get(c *gin.Context) {
	eventID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not found"})
		return
	}

	service := services.NewEventService(repositories.NewEventRepository(db.(*sql.DB)))
	event, err := service.Find(eventID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (e EventHandler) Create(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database not found"})
		return
	}

	service := services.NewEventService(repositories.NewEventRepository(db.(*sql.DB)))
	if err := service.Create(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, event)
}

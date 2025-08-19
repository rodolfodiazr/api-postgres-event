package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" binding:"required,max=100"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required,gtfield=StartTime"`
	CreatedAt   time.Time `json:"created_at"`
}

type Events []Event

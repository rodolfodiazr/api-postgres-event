package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/rodolfodiazr/api-postgres-event/models"
)

type EventRepository interface {
	Create(event *models.Event) error
	Find(id uuid.UUID) (models.Event, error)
	List() (models.Events, error)
}

type DefaultEventRepository struct {
	db *sql.DB
}

func (r *DefaultEventRepository) Create(event *models.Event) error {
	query := `
		INSERT INTO events (title, description, start_time, end_time)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	return r.db.QueryRow(query, event.Title, event.Description, event.StartTime, event.EndTime).
		Scan(&event.ID, &event.CreatedAt)
}

func (r *DefaultEventRepository) Find(id uuid.UUID) (models.Event, error) {
	var event models.Event
	query := `
		SELECT id, title, description, start_time, end_time, created_at 
		FROM events 
		WHERE id=$1
	`
	err := r.db.QueryRow(query, id).Scan(&event.ID, &event.Title, &event.Description, &event.StartTime, &event.EndTime, &event.CreatedAt)
	return event, err
}

func (r *DefaultEventRepository) List() (models.Events, error) {
	events := models.Events{}
	query := `SELECT id, title, description, start_time, end_time, created_at FROM events ORDER BY start_time ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var e models.Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.StartTime, &e.EndTime, &e.CreatedAt); err != nil {
			return events, err
		}
		events = append(events, e)
	}
	return events, nil
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &DefaultEventRepository{db: db}
}

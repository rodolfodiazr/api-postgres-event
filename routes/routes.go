package routes

import (
	"encoding/json"
	"fmt"
	"ltk/internal"
	"ltk/models"
	"net/http"
	"time"
)

var mux = http.NewServeMux()

func Register() *http.ServeMux {
	mux.Handle("/events", getEvents())
	return mux
}

func getEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Get(w, r)
		case http.MethodPost:
			Create(w, r)
		default:
			json.NewEncoder(w).Encode("Not Found 404")
		}
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	events := []models.Event{
		{Title: "Any title"},
	}
	json.NewEncoder(w).Encode(events)
}

func Create(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	db, err := internal.InitDB()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	query := `
        INSERT INTO events (title, description, start_time, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
	err = db.QueryRow(query, event.Title, event.Description, event.StartTime, time.Now()).
		Scan(&event.ID, event.CreatedAt)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("Event %v created", event.ID))
}

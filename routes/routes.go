package routes

import (
	"encoding/json"
	"fmt"
	"ltk/internal"
	"ltk/models"
	"net/http"
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
		fmt.Println("err: ", err)
		return
	}

	db, err := internal.InitDB()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	db.Create(&event)
}

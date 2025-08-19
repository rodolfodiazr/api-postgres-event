package services

import (
	"github.com/google/uuid"
	"github.com/rodolfodiazr/api-postgres-event/models"
	"github.com/rodolfodiazr/api-postgres-event/repositories"
)

type EventService interface {
	Create(event *models.Event) error
	Find(id uuid.UUID) (models.Event, error)
	List() (models.Events, error)
}

type DefaultEventService struct {
	repo repositories.EventRepository
}

func (r *DefaultEventService) Create(event *models.Event) error {
	return r.repo.Create(event)
}

func (r *DefaultEventService) Find(id uuid.UUID) (models.Event, error) {
	return r.repo.Find(id)
}

func (r *DefaultEventService) List() (models.Events, error) {
	return r.repo.List()
}

func NewEventService(repo repositories.EventRepository) EventService {
	return &DefaultEventService{repo: repo}
}

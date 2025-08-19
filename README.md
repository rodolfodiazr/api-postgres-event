# Objective: Build a RESTful API in Go with PostgreSQL to manage Events.

## Event Model
Include:
* id (UUID, auto-generated)
* title (string, required, max 100 characters)
* description (string, optional)
* start_time, end_time (timestamps; start_time < end_time)
* created_at (timestamp, auto-set on creation)

## API Endpoints
POST /events
* Accept JSON input
* Validate fields and time constraints
* Persist to DB and return created record with HTTP 201

GET /events
* Return all events ordered by start_time ascending

GET /events/{id}`
* Return event by UUID or HTTP 404 if not found

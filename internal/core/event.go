package core

import "time"

// Event represents something that happened in the system.
type Event struct {
	Name    string                 // e.g. "purchase", "login"
	ActorID string                 // player/user id
	At      time.Time              // event time
	Attrs   map[string]any         // flexible payload
	TraceID string                 // optional tracing id
}

package core

// Trigger builds actions when a rule matches.
type Trigger func(Event) []Action

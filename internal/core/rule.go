package core

// Predicate decides whether a rule matches an event.
type Predicate func(Event) bool

// Rule describes a LiveOps rule.
type Rule struct {
	ID          string
	Description string
	When        Predicate
	Triggers    []Trigger
	Enabled     bool
}

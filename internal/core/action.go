package core

// Action represents an effect to be executed by adapters (not by core).
type Action struct {
	Type   string                 // e.g. "grant_item"
	Target string                 // usually ActorID
	Params map[string]any         // action parameters
}

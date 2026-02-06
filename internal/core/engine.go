package core

// Engine evaluates events against rules and returns actions.
type Engine struct {
	rules []Rule
}

func NewEngine(rules []Rule) *Engine {
	cp := make([]Rule, 0, len(rules))
	for _, r := range rules {
		cp = append(cp, r)
	}
	return &Engine{rules: cp}
}

func (e *Engine) Evaluate(ev Event) []Action {
	var out []Action

	for _, r := range e.rules {
		if !r.Enabled || r.When == nil {
			continue
		}
		if !r.When(ev) {
			continue
		}
		for _, t := range r.Triggers {
			if t == nil {
				continue
			}
			out = append(out, t(ev)...)
		}
	}
	return out
}

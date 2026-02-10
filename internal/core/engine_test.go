package core

import (
	"testing"
	"time"
)

func TestEngine_Evaluate_PurchaseRule(t *testing.T) {
	grantItem := func(itemID string) Trigger {
		return func(e Event) []Action {
			return []Action{
				{
					Type:   "grant_item",
					Target: e.ActorID,
					Params: map[string]any{"item_id": itemID},
				},
			}
		}
	}

	isPurchaseOver := func(min int) Predicate {
		return func(e Event) bool {
			if e.Name != "purchase" {
				return false
			}
			v, ok := e.Attrs["amount"]
			if !ok {
				return false
			}
			amt, ok := v.(int)
			return ok && amt >= min
		}
	}

	rules := []Rule{
		{
			ID:          "purchase_100",
			Description: "purchase >= 100 grants starter pack",
			When:        isPurchaseOver(100),
			Triggers:    []Trigger{grantItem("starter_pack")},
			Enabled:     true,
		},
	}

	engine := NewEngine(rules)

	ev := Event{
		Name:    "purchase",
		ActorID: "player-1",
		At:      time.Now(),
		Attrs:   map[string]any{"amount": 120},
	}

	actions := engine.Evaluate(ev)
	if len(actions) != 1 {
		t.Fatalf("expected 1 action, got %d", len(actions))
	}
	if actions[0].Type != "grant_item" {
		t.Fatalf("unexpected action type: %s", actions[0].Type)
	}
}

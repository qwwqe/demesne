package game

// EndCondition describes end conditions for a game.
type EndCondition interface {
	// Evaluate determines whether end conditions have been met.
	Evaluate(game) bool
}

// EmptySupplyEndCondition is an end condition that looks at the
// number of empty piles in the supply.
type EmptySupplyEndCondition struct {
	Count int
}

// Evaluate determines whether the end condition has been met.
func (ec EmptySupplyEndCondition) Evaluate(g game) bool {
	empty := 0

	for _, pile := range g.Supply {
		if pile.Size() == 0 {
			empty++
		}

		if empty >= ec.Count {
			return true
		}
	}

	return false
}

var _ EndCondition = EmptySupplyEndCondition{}

// EmptyPileEndCondition is an end condition that looks at the
// size of a specific pile.
type EmptyPileEndCondition struct {
	Id string
}

// Evaluate determines whether the end condition has been met.
func (ec EmptyPileEndCondition) Evaluate(g game) bool {
	// NOTE: Should a Game be maintaining a lookup so that
	// checks like this don't need to keep iterating over
	// all the piles?
	for _, pile := range g.Supply {
		if pile.Id == ec.Id {
			return pile.Size() == 0
		}
	}

	return false
}

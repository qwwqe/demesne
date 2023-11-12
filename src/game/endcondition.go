package game

// EndCondition describes end conditions for a game.
//
// NOTE: See comment below about converting this to a struct
// that exposes an Evaluate() method
type EndCondition interface {
	// Evaluate determines whether end conditions have been met.
	Evaluate(Game) bool
}

type EndConditionSpec struct {
	EndCondition
}

func (s EndConditionSpec) Build() EndCondition {
	return s
}

// EmptySupplyEndCondition is an end condition that looks at the
// number of empty piles in the supply.
//
// NOTE: The structure of an EndCondition probably needs to be more
// formally (and generically) defined, or dealing with various
// "types" of end conditions will likely be a headache.
// A more ergonomic approach might be to have a single struct
// that can be used to describe arbitrary end conditions,
// and then a few convenience functions that return structs
// satisfying common conditions.
type EmptySupplyEndCondition struct {
	Count int
}

// Evaluate determines whether the end condition has been met.
func (ec EmptySupplyEndCondition) Evaluate(g Game) bool {
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
func (ec EmptyPileEndCondition) Evaluate(g Game) bool {
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

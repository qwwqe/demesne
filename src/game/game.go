package game

// A game of Demesne.
type game struct {
	// Id uniquely identifies a Game.
	Id string

	// The current turn.
	//
	// NOTE: This state information is actually specific
	// to a certain stage of the game, and should probably
	// be stored within that stage, rather than within the
	// game itself.
	Turn int

	// The current stage.
	//
	// TODO: Evaluate necessity of explicit stages in a game.
	Stage Stage

	// The Players in the Game.
	//
	// The order of this array determines turn playing order.
	// In other words, the current player is given by the following relation:
	// 		currentPlayer := Players[Turn % len(Players)]
	Players []Player

	// The Trash.
	Trash Pile

	// The Supply is the collection of all card Piles which can be
	// directly purchased from in a given game of Demesne.
	Supply []SupplyPile

	EndConditions []EndCondition
}

func (g game) IsFinished() bool {
	for _, condition := range g.EndConditions {
		if condition.Evaluate(g) {
			return true
		}
	}

	for _, pile := range g.Supply {
		for _, condition := range pile.EndConditions {
			if condition.Evaluate(g) {
				return true
			}
		}
	}

	return false
}

// A Stage is a distinct state in the life cycle of a game.
//
// NOTE: Stub.
// NOTE: It may not be necessary to categorize the game into stages here.
type Stage interface{}

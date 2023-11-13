package game

// A game of Demesne.
type Game struct {
	// Id uniquely identifies a game.
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

	// The Players in the game.
	//
	// The order of this array determines turn playing order.
	// In other words, the current player is given by the following relation:
	// 		currentPlayer := Players[Turn % len(Players)]
	Players []Player

	// The Trash.
	Trash Pile

	// The Supply is the collection of all card Piles which can be
	// directly purchased from in a given game of Demesne.
	Supply []Pile

	// EndConditions define the end conditions for the game.
	EndConditions []EndCondition

	// DealRules define how cards are dealt.
	//
	// TODO: Move this to GameSpec.
	DealRules []DealRule

	// RandomSeed is the seed used for all randomness in the game.
	RandomSeed int64
}

// IsFinished returns a boolean value representing whether the
// game has satisfied the end conditions described in the rule set.
func (g Game) IsFinished() bool {
	for _, condition := range g.EndConditions {
		if condition.Evaluate(g) {
			return true
		}
	}

	return false
}

// A Stage is a distinct state in the life cycle of a game.
//
// NOTE: Stub.
// NOTE: It may not be necessary to categorize the game into stages here.
type Stage interface{}

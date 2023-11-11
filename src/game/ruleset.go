package game

import (
	"github.com/google/uuid"
)

// A GameSpec defines the rules for creating the supply, dealing player decks,
// and determining end of game.
type GameSpec struct {
	// Piles defined as being in the Supply.
	SupplySpec []SupplyPileSpec

	// Predicates determining completion of the game.
	// Game end conditions may also be specified by individual CardSets.
	//
	// The set of end conditions are evaluated as a logical union,
	// meaning that if any are true, the game as a whole is
	// judged to be over.
	EndConditions []EndCondition
}

func (r GameSpec) Build(numPlayers int) Game {
	g := Game{}

	g.Players = make([]Player, numPlayers)
	for _, player := range g.Players {
		player.Id = uuid.NewString()
	}

	g.Supply = make([]Pile, len(r.SupplySpec))
	for i, pileSpec := range r.SupplySpec {
		g.Supply[i] = pileSpec.Build(numPlayers)
		for _, player := range g.Players {
			player.Deck.AddCards(g.Supply[i].Deal())
		}
	}

	return g
}

// SupplyPileSpec defines how a supply pile is created,
// how it contributes to the initial deal,
// and how it influences the end of game conditions.
type SupplyPileSpec struct {
	EndConditions []EndCondition
	DealRules     []DealRule
	PileSpec      []PileSpec
}

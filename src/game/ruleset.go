package game

import (
	"github.com/google/uuid"
	"github.com/qwwqe/demesne/src/card"
)

// A RuleSet defines the rules for creating the supply, dealing player decks,
// and determining end of game.
//
// NOTE: Do Kingdom Sets and Base Sets really need to
// be separated structurally? It might be simpler to just
// use properties of the CardSet itself to determine this.
type RuleSet struct {
	// Piles defined as being in the Supply.
	SupplySpec []PileSpec

	// Predicates determining completion of the game.
	// Game end conditions may also be specified by individual CardSets.
	//
	// The set of end conditions are evaluated as a logical union,
	// meaning that if any are true, the game as a whole is
	// judged to be over.
	EndConditions []EndCondition
}

func (r RuleSet) BuildGame(numPlayers int) game {
	g := game{}

	g.Players = make([]Player, numPlayers)
	for _, player := range g.Players {
		player.Id = uuid.NewString()
	}

	g.Supply = make([]SupplyPile, len(r.SupplySpec))
	for i, pileSpec := range r.SupplySpec {
		g.Supply[i] = pileSpec.Build(numPlayers)
		for _, player := range g.Players {
			player.Deck.AddCards(g.Supply[i].Deal())
		}
	}

	return g
}

// PileSpec defines how a pile is created, how it contributes to the initial deal,
// and how it influences the end of game conditions.
//
// NOTE: Does this really need to be an interface?
//
// NOTE: Can Deal() and EndConditions() be offloaded onto Pile instead?
type PileSpec interface {
	Build(numPlayers int) card.Pile
	Deal(pile *card.Pile) []card.Card
	EndConditions() []EndCondition
}

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
	SupplySpec

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

	// TODO: Combine BaseCard and KingdomCard so the following doesn't need
	// to be repeated twice whenever we want to iterate over all the cards in the supply.
	for _, basePile := range r.BasePileSpecs {
		g.BaseCards = append(g.BaseCards, basePile.Build(numPlayers))
		for _, player := range g.Players {
			player.Deck.AddCards(basePile.Deal(&g.BaseCards[len(g.BaseCards)-1]))
		}
	}

	for _, kingdomPile := range r.KingdomPileSpecs {
		g.KingdomCards = append(g.KingdomCards, kingdomPile.Build(numPlayers))
		for _, player := range g.Players {
			player.Deck.AddCards(kingdomPile.Deal(&g.KingdomCards[len(g.KingdomCards)-1]))
		}
	}

	return g
}

// IsGameFinished returns a boolean value representing whether the
// game has satisfied the end conditions described in the rule set.
func (r RuleSet) IsGameFinished(g game) bool {
	for _, condition := range r.EndConditions {
		if condition.Evaluate(g) {
			return true
		}
	}

	for _, cardSet := range r.SupplySpec.All() {
		for _, condition := range cardSet.EndConditions() {
			if condition.Evaluate(g) {
				return true
			}
		}
	}

	return false
}

type SupplySpec struct {
	BasePileSpecs    []PileSpec
	KingdomPileSpecs []PileSpec
}

// All is a convenience function for iterating over
// card sets in a supply set.
//
// TODO: Make this an actual iterator when 1.22 lands.
func (s SupplySpec) All() []PileSpec {
	specs := make([]PileSpec, 0, len(s.BasePileSpecs)+len(s.KingdomPileSpecs))

	specs = append(specs, s.BasePileSpecs...)
	specs = append(specs, s.KingdomPileSpecs...)

	return specs
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

package game

import (
	"github.com/google/uuid"
	"github.com/qwwqe/demesne/src/card"
)

type EndCondition func(game) bool

// A RuleSet defines the rules for creating the supply, dealing player decks,
// and determining end of game.
//
// NOTE: Do Kingdom Sets and Base Sets really need to
// be separated structurally? It might be simpler to just
// use properties of the CardSet itself to determine this.
type RuleSet struct {
	// Card sets defined as being in the Supply.
	SupplySet

	// Predicates determining completion of the game.
	// Game end conditions may also be specified by individual CardSets.
	//
	// The set of end conditions are evaluated as a logical union,
	// meaning that if any are true, the game as a whole is
	// judged to be over.
	EndConditions []EndCondition
}

func (r RuleSet) buildSupplyPile(cardSet CardSet) card.Pile {
	// NOTE: Maybe these properties should be specified in the RuleSet?
	pile := card.Pile{
		Countable:  true,
		Faceup:     true,
		Browseable: false,
	}

	card := cardSet.Card()
	amount := cardSet.PileSize(len(g.Players))
	for i := 0; i < amount; i++ {
		pile.AddCard(card.Clone())
	}

	return pile
}

func (r RuleSet) BuildGame(players int) game {
	g := game{}

	g.Players = make([]Player, players)
	for _, player := range g.Players {
		player.Id = uuid.NewString()
	}

	setToPileMap := map[string]*card.Pile{}

	for _, bs := range r.BaseCardSets {
		g.BaseCards = append(g.BaseCards, r.buildSupplyPile(bs))
		setToPileMap[bs.Card().Name] = &g.BaseCards[len(g.BaseCards)-1]
	}

	for _, ks := range r.KingdomCardSets {
		g.KingdomCards = append(g.KingdomCards, r.buildSupplyPile(ks))
		setToPileMap[ks.Card().Name] = &g.KingdomCards[len(g.KingdomCards)-1]
	}

	for _, cardSet := range r.SupplySet.All() {
		for _, player := range g.Players {
			amount, deductFromPile := cardSet.DealAmount()
			for i := 0; i < amount; i++ {
				player.Deck.AddCard(cardSet.Card())
				if deductFromPile {
					// setToPileMap
					// TODO: deal with pile adjustments.
				}
			}
		}
	}

	return g
}

// func (r RuleSet)

// IsGameFinished returns a boolean value representing whether the
// game has satisfied the end conditions described in the rule set.
func (r RuleSet) IsGameFinished(g game) bool {
	for _, condition := range r.EndConditions {
		if condition(g) {
			return true
		}
	}

	for _, cardSet := range r.SupplySet.All() {
		for _, condition := range cardSet.EndConditions() {
			if condition(g) {
				return true
			}
		}
	}

	return false
}

type SupplySet struct {
	BaseCardSets    []CardSet
	KingdomCardSets []CardSet
}

// All is a convenience function for iterating over
// card sets in a supply set.
//
// TODO: Make this an actual iterator when 1.22 lands.
func (s SupplySet) All() []CardSet {
	sets := make([]CardSet, 0, len(s.BaseCardSets)+len(s.KingdomCardSets))

	sets = append(sets, s.BaseCardSets...)
	sets = append(sets, s.KingdomCardSets...)

	return sets
}

type CardSet interface {
	Card() card.Card
	PileSize(players int) int
	DealAmount() (amount int, deductFromPile bool)
	EndConditions() []EndCondition
}

/**
 * Reference material below.
 *
 * TODO: Turn these into test cases or something.
 *
 */

// A simple end condition based on supply pile exhaustion.
//
// This is mostly intended for reference until a more comprehensive
// framework for dynamic definition is established.
func BasicSupplyEndCondition(g game) bool {
	supplyPilesExhausted := 0
	for _, p := range g.Supply.BaseCards {
		if p.Size() == 0 {
			supplyPilesExhausted++
		}
	}

	for _, p := range g.Supply.KingdomCards {
		if p.Size() == 0 {
			supplyPilesExhausted++
		}
	}

	if len(g.Players) < 5 {
		return supplyPilesExhausted >= 3
	}

	return supplyPilesExhausted >= 4
}

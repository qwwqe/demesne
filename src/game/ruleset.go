package game

import "github.com/qwwqe/demesne/src/card"

// A RuleSet defines the rules for creating the supply, dealing player decks,
// and determining end of game.
//
// NOTE: Do Kingdom Sets and Base Sets really need to
// be separated structurally? It might be simpler to just
// use properties of the CardSet itself to deermine this.
type RuleSet struct {
	// Card sets defined as being in the Supply.
	SupplySet

	// Predicates determining completion of the game.
	//
	// The set of end conditions are evaluated as a logical union,
	// meaning that if any are true, the game as a whole is
	// judged to be over.
	EndConditions []func(game) bool
}

func (r RuleSet) SetupTable(g *game) {
	// NOTE: It's not clear whether table setup should really
	// be handling this kind of cleanup
	g.BaseCards = g.BaseCards[:0]
	g.KingdomCards = g.KingdomCards[:0]
	for _, p := range g.Players {
		p.Hand = p.Hand[:0]
		p.PlayArea = p.PlayArea[:0]
		p.Deck.Drain()
		p.Discard.Drain()
	}

	buildSupplyPile := func(cardSet CardSet) card.Pile {
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

	for _, bs := range r.BaseCardSets {
		g.BaseCards = append(g.BaseCards, buildSupplyPile(bs))
	}

	for _, ks := range r.KingdomCardSets {
		g.KingdomCards = append(g.KingdomCards, buildSupplyPile(ks))
	}

	// TODO: Deal decks.
}

// IsGameFinished returns a boolean value representing whether the
// game has satisfied the end conditions described in the rule set.
func (r RuleSet) IsGameFinished(g game) bool {
	for _, condition := range r.EndConditions {
		if condition(g) {
			return true
		}
	}

	return false
}

type SupplySet struct {
	BaseCardSets    []CardSet
	KingdomCardSets []CardSet
}

type CardSet interface {
	Card() card.Card
	PileSize(players int) int
	// IsGameFinished(game) bool
}

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

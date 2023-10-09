package game

import "github.com/qwwqe/demesne/src/card"

// A RuleSet defines how a game of Demesne is set up
// and how end game is determined, among other things.
//
// NOTE: Do Kingdom Sets and Base Sets really need to
// be separated structurally? It might be simpler to just
// use properties of the CardSet itself to deermine this.
type RuleSet struct {
	SupplySet
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
		amount := cardSet.Amount(len(g.Players))
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

// Determine whether game end conditions have been met.
//
// TODO: Figure out a way to make this more flexible instead of
// hard-coding the three-pile logic into the rule set struct receiver itself.
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
	Amount(players int) int
	// IsGameFinished(game) bool
}

func BaseGameEndCondition(g game) bool {
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

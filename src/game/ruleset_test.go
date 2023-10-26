package game

import (
	"github.com/qwwqe/demesne/src/card"
)

// A simple end condition based on supply pile exhaustion.
//
// This is mostly intended for reference until a more comprehensive
// framework for dynamic definition is established.
func basicSupplyEndCondition(g game) bool {
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

type provinceCardSet struct{}

func (cs provinceCardSet) Card() card.Card {
	return card.Card{
		Name: "province",
	}
}

func (cs provinceCardSet) BuildPile(numPlayers int) card.Pile {
	pileSize := 12
	if numPlayers == 2 {
		pileSize = 8
	} else if numPlayers == 5 {
		pileSize = 15
	} else if numPlayers == 6 {
		pileSize = 18
	}

	pile := card.Pile{
		Countable:  true,
		Faceup:     true,
		Browseable: false,
	}

	// NOTE: See note for CardSet.Card()
	card := cs.Card()
	for i := 0; i < pileSize; i++ {
		pile.AddCard(card.Clone())
	}

	return pile
}

func (cs provinceCardSet) EndConditions() []endCondition {
	return []endCondition{
		// End condition for when the Province pile is emptied.
		//
		// TODO: Find a better way of mapping card sets to piles.
		func(g game) bool {
			found := false
			for _, pile := range g.Supply.All() {
				if pile.Size() > 0 && pile.Cards[0].Name == cs.Card().Name {
					found = true
					break
				}
			}

			return found
		},
	}
}

func (cs provinceCardSet) Deal(*card.Pile) []card.Card {
	return nil
}

var _ CardSet = provinceCardSet{}

type estateCardSet struct{}

func (cs estateCardSet) Card() card.Card {
	return card.Card{
		Name: "estate",
	}
}

func (cs estateCardSet) BuildPile(numPlayers int) card.Pile {
	pileSize := 8
	if numPlayers != 2 {
		pileSize = 12
	}

	amountPerPlayer := 3
	pileSize += amountPerPlayer * numPlayers

	pile := card.Pile{
		Countable:  true,
		Faceup:     true,
		Browseable: false,
	}

	// NOTE: See note for CardSet.Card()
	card := cs.Card()
	for i := 0; i < pileSize; i++ {
		pile.AddCard(card.Clone())
	}

	return pile
}

func (cs estateCardSet) Deal(pile *card.Pile) []card.Card {
	amountPerPlayer := 3
	return pile.Draw(amountPerPlayer)
}

func (cs estateCardSet) EndConditions() []endCondition {
	return []endCondition{}
}

var _ CardSet = estateCardSet{}

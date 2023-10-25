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

func (r RuleSet) BuildGame(numPlayers int) game {
	g := game{}

	g.Players = make([]Player, numPlayers)
	for _, player := range g.Players {
		player.Id = uuid.NewString()
	}

	// TODO: Combine BaseCard and KingdomCard so the following doesn't need
	// to be repeated twice whenever we want to iterate over all the cards in the supply.
	for _, bs := range r.BaseCardSets {
		g.BaseCards = append(g.BaseCards, bs.BuildPile(numPlayers))
		for _, player := range g.Players {
			player.Deck.AddCards(bs.Deal(&g.BaseCards[len(g.BaseCards)-1]))
		}
	}

	for _, ks := range r.KingdomCardSets {
		g.KingdomCards = append(g.KingdomCards, ks.BuildPile(numPlayers))
		for _, player := range g.Players {
			player.Deck.AddCards(ks.Deal(&g.KingdomCards[len(g.KingdomCards)-1]))
		}
	}

	return g
}

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

// NOTE: Use a withPlayers() instead of passing the number of players
// to every method? A little less clean in terms of function application,
// but it would provide better guarantees about consistency between
// behaviour that depends on a fixed player count (like building piles and dealing).
type CardSet interface {
	// NOTE: Realistically, if the CardSet exposes both a BuildPile and DealCards
	// method, there isn't really any reason to expose a Card method as well.
	// It also ceases being useful when split piles are introduced.
	Card() card.Card
	BuildPile(numPlayers int) card.Pile
	Deal(pile *card.Pile) []card.Card
	// PileSize(players int) int
	// DealAmount() (amount int, deductFromPile bool)
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

func (cs provinceCardSet) EndConditions() []EndCondition {
	return []EndCondition{
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

func (cs estateCardSet) EndConditions() []EndCondition {
	return []EndCondition{}
}

var _ CardSet = estateCardSet{}

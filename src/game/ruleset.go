package game

import (
	"github.com/google/uuid"
	"github.com/qwwqe/demesne/src/card"
)

type endCondition func(game) bool

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
	EndConditions []endCondition
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
	EndConditions() []endCondition
}

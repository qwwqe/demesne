package game

import "github.com/qwwqe/demesne/src/card"

// A RuleSet defines how a game of Demesne is set up
// and how end game is determined, among other things.
type RuleSet struct {
	CardSets []CardSet
}

type CardSet interface {
	Card() card.Card
	Amount(players int) int
}

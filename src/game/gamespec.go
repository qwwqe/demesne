package game

import (
	"time"

	"github.com/google/uuid"
)

// A GameSpec defines the rules for creating the supply, dealing player decks,
// and determining end of game.
type GameSpec struct {
	// Piles defined as being in the Supply.
	SupplyPileSpecs []SupplyPileSpec

	// Predicates determining completion of the game.
	// Game end conditions may also be specified by individual CardSets.
	//
	// The set of end conditions are evaluated as a logical union,
	// meaning that if any are true, the game as a whole is
	// judged to be over.
	EndConditionSpecs []EndConditionSpec
}

func (gs GameSpec) Validate() bool {
	return true
}

func (gs GameSpec) Build(numPlayers int) Game {
	// TODO: Validate first?

	g := Game{}

	g.Id = uuid.NewString()

	g.RandomSeed = time.Now().UnixNano()

	g.Players = make([]Player, numPlayers)
	for _, player := range g.Players {
		player.Id = uuid.NewString()
		player.Deck.randomSeed = g.RandomSeed
	}

	g.Supply = make([]Pile, 0, len(gs.SupplyPileSpecs))
	for _, supplyPileSpec := range gs.SupplyPileSpecs {
		// Build pile
		g.Supply = append(g.Supply, supplyPileSpec.PileSpec.Build(numPlayers))
		pile := &g.Supply[len(g.Supply)-1]

		// Deal from pile
		for _, dealRuleSpec := range supplyPileSpec.DealRuleSpecs {
			dealRule := dealRuleSpec.Build()
			for _, player := range g.Players {
				player.Deck.AddCards(dealRule.Deal(pile))
			}
		}

		// Register pile-stipulated end conditions
		for _, endConditionSpec := range supplyPileSpec.EndConditionSpecs {
			g.EndConditions = append(g.EndConditions, endConditionSpec.Build())
		}
	}

	for _, player := range g.Players {
		player.Deck.Shuffle()
	}

	// Register general end conditions
	for _, endConditionSpec := range gs.EndConditionSpecs {
		g.EndConditions = append(g.EndConditions, endConditionSpec)
	}

	return g
}

// SupplyPileSpec defines how a supply pile is created,
// how it contributes to the initial deal,
// and how it influences the end of game conditions.
type SupplyPileSpec struct {
	EndConditionSpecs []EndConditionSpec
	DealRuleSpecs     []DealRuleSpec
	PileSpec          PileSpec
}

// func (s SupplyPileSpec) Build(numPlayers int) ([]EndCondition, []DealRule, Pile) {

// }

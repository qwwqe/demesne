package game

import (
	"github.com/google/uuid"
)

// A Card represents a particular, concrete instance of a Demesne card.
type Card struct {
	// Id identifies a specific and concrete Card, such as one being held by a
	// particular Player or one currently in a Player's play area.
	// Two Card instances with the same Id shall be viewed as equivalent
	// for all purposes.
	Id string

	// Name identifies the abstract definition of a Card.
	// Any Card instances possessing the same Name should also possess
	// identical properties, not including Id.
	Name string

	// Cost is the set of requirements that must be met to purchase this card.
	Cost Cost

	// Types.
	Types []CardType

	ActionEffects []Effect

	ReactionEffects []Reaction

	TreasureEffects []Effect

	ScoringEffects []Effect
}

// Return a deep copy of the given Card, with Id set to a random UUID.
func (c Card) Clone() Card {
	nc := c
	nc.Id = uuid.NewString()
	return nc
}

type CardType string

const (
	CardTypeAction   CardType = "action"
	CardTypeReaction CardType = "reaction"
	CardTypeTreasure CardType = "treasure"
	CardTypeAttack   CardType = "attack"
	CardTypeVictory  CardType = "victory"
	CardTypeCurse    CardType = "curse"
)

// A CardSpec defines a Card.
type CardSpec struct {
	// Name identifies a CardSpec.
	Name string `yaml:"name"`

	// Cost defines the requirements needed to purchase the card specified.
	Cost CostSpec `yaml:"cost"`

	// Types defines the types of the card specified.
	Types []CardTypeSpec `yaml:"types"`

	// ActionEffects define what a card does when played as an action.
	ActionEffects []EffectSpec `yaml:"actionEffects"`

	// ReactionEffects define what a card does when played as a reaction.
	ReactionEffects []ReactionSpec `yaml:"reactionEffects"`

	// ScoringEffects define what a card does when evluated during end-game scoring.
	ScoringEffects []EffectSpec `yaml:"scoringEffects"`
}

func (s CardSpec) Build() Card {
	// Still no slice map? ...?
	types := []CardType{}
	for _, t := range s.Types {
		types = append(types, t.Build())
	}

	actionEffects := []Effect{}
	for _, e := range s.ActionEffects {
		actionEffects = append(actionEffects, e.Build())
	}

	reactionEffects := []Reaction{}
	for _, r := range s.ReactionEffects {
		reactionEffects = append(reactionEffects, r.Build())
	}

	scoringEffects := []Effect{}
	for _, s := range s.ScoringEffects {
		scoringEffects = append(scoringEffects, s.Build())
	}

	return Card{
		Id:              uuid.NewString(),
		Name:            s.Name,
		Cost:            s.Cost.Build(),
		Types:           types,
		ActionEffects:   actionEffects,
		ReactionEffects: reactionEffects,
		ScoringEffects:  scoringEffects,
	}
}

type CardTypeSpec CardType

func (s CardTypeSpec) Build() CardType {
	return CardType(s)
}

// Cost represents the set of requirements that must be met to purchase a card.
type Cost struct {
	// The fixed treasure cost for this card.
	Treasure AmountFixed
}

type CostSpec struct {
	Treasure AmountFixed `yaml:"treasure"`
}

func (s CostSpec) Build() Cost {
	return Cost{
		Treasure: s.Treasure,
	}
}

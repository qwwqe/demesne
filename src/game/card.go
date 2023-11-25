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

	// Effects.
	Effects []Effect
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
}

func (s CardSpec) Build() Card {
	return Card{
		Id:   uuid.NewString(),
		Name: s.Name,
	}
}

type CardTypeSpec CardType

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

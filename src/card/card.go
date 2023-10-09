package card

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// A Card simltaneously represents the definition of a Demesne card and
// a particular, concrete instance of such definition.
//
// NOTE: When implementing functionality involving properties of concrete
// cards (such as tokens, counters or even other cards being placed on top of
// them), it may be more ergonomic to separate the definition and instance
// of a card.
type Card struct {
	// Id identifies a specific and concrete Card, such as one being held by a
	// particular Player or one currently in a Player's play area.
	// Two Card instances with the same Id shall be viewed as equivalent
	// for all purposes.
	Id string

	// Name identifies the abstract definition of a Card.
	// Any Card instances possessing the same Name should also possess
	// properties except for Id.
	//
	// NOTE: As mentioned above, it may be a more flexible design to
	// implement via separation of card definitions and insances.
	Name string
}

// Return a deep copy of the given Card, with Id set to a random UUID.
func (c Card) Clone() Card {
	nc := c
	nc.Id = uuid.NewString()
	return nc
}

// A Pile is an ordered collection of Cards.
//
// NOTE: Does a Pile need an Id?
type Pile struct {
	// The cards that constitute this Pile.
	Cards []Card

	// Whether the first card in this Pile is visible to all players.
	Faceup bool

	// Whether the number of cards in this Pile is known to all players.
	//
	// NOTE: In certain cases it is disputed whether the size of a Pile
	// can be known by its owner but not by other players (i.e. a player's Deck).
	// This may require specification of a finer grain, such as by replacing
	// this field with CountableByOthers and CountableByOwner.
	Countable bool

	// Whether the cards in this Pile may be inspected at will by all players.
	Browseable bool
}

// Stub.
func (p Pile) Top() *Card {
	if len(p.Cards) == 0 {
		return nil
	}

	return &p.Cards[0]
}

// Stub.
func (p *Pile) AddCard(c Card) *Pile {
	p.Cards = append(p.Cards, c)
	return p
}

// Stub.
func (p *Pile) AddCards(cs []Card) *Pile {
	p.Cards = append(p.Cards, cs...)
	return p
}

// Stub.
func (p *Pile) Shuffle() *Pile {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(p.Cards), func(i, j int) {
		p.Cards[i], p.Cards[j] = p.Cards[j], p.Cards[i]
	})

	return p
}

// Stub.
func (p *Pile) Draw(n int) []Card {
	if len(p.Cards) < n || n == 0 {
		return nil
	}

	cs := p.Cards[0:n]
	p.Cards = p.Cards[n:]
	return cs
}

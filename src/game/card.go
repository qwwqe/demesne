package game

import (
	"math/rand"
	"time"

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
}

// Return a deep copy of the given Card, with Id set to a random UUID.
func (c Card) Clone() Card {
	nc := c
	nc.Id = uuid.NewString()
	return nc
}

// A CardSpec defines a Card.
type CardSpec struct {
	// Name identifies a CardSpec.
	Name string
}

// A Pile is an ordered collection of Cards.
//
// NOTE: Should a Pile contain information about dealing and end of game conditions?
// That is to say, should it actually expose the Deal() and EndCondition()
// methods defined in PileSpec?
type Pile struct {
	// Id uniquely identifies a specific and concrete Pile.
	Id string

	// Name identifies the abstract definition of a Pile.
	Name string

	// Kind describes what role a Pile takes on during a game.
	Kind PileKind

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

type PileKind string

const (
	KingdomPile PileKind = "kingdom"
	BasePile    PileKind = "base"
	GenericPile PileKind = "generic"
)

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

func (p *Pile) Drain() []Card {
	cs := p.Cards[:]
	p.Cards = p.Cards[:0]
	return cs
}

func (p Pile) Size() int {
	return len(p.Cards)
}

// PileSpec defines how a pile is created.
type PileSpec struct {
	CardSpec CardSpec

	PileSizeSpec PileSizeSpec
}

// PileSpecSpec describes how the size of a pile is determined.
//
// NOTE: This struct will probably need to be re-visited when new
// constraints for defining pile size are introduced.
type PileSizeSpec struct {
	// DefaultPileSize defines the default pile size.
	DefaultPileSize int

	// PlayerCountPileSizeSpecs describes pile size specifications
	// based on player count.
	//
	// TODO: Revisit the naming for this...
	PlayerCountPileSizeSpecs []PlayerCountPileSizeSpec
}

// PlayerCountPileSizeSpec describes how the size of a pile relates
// to player count.
type PlayerCountPileSizeSpec struct {
	PlayerCount int
	PileSize    int
}

// A SupplyPile is a Pile which also contains information about
// dealing cards and game end conditions.
//
// NOTE: Does this really need to be separate from Pile?
// Is this going to cause headaches down the road?
//
// NOTE: Can Dealer be moved into the Game model instead?
// Then responsibilities for dealing cards wouldn't be awkwardly
// offloaded onto piles in the supply.
//
// NOTE: Can EndConditions be moved into the Game model as well?
// Maybe encapsulated in some kind of Arbiter model?
type SupplyPile struct {
	Pile

	EndConditions []EndCondition

	Dealer Dealer
}

func (sp *SupplyPile) Deal() []Card {
	return sp.Dealer.Deal(&sp.Pile)
}

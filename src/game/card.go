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

	// Cost is the set of requirements that must be met to purchase this card.
	Cost Cost
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

func (s CardSpec) Build() Card {
	return Card{
		Id:   uuid.NewString(),
		Name: s.Name,
	}
}

// Cost represents the set of requirements that must be met to purchase a card.
type Cost struct {
	Treasure FixedAmount
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

	// Random seed used for things like shuffling.
	// If this is equal to the zero, the current time will instead be used.
	//
	// TODO: Make this be a thread-safe source instead of generating
	// the source every time.
	randomSeed int64
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
	seed := p.randomSeed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	r := rand.New(rand.NewSource(seed))

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

func (ps PileSpec) Build(numPlayers int) Pile {
	pileSize := ps.PileSizeSpec.Build(numPlayers)

	pile := Pile{
		Countable:  true,
		Faceup:     true,
		Browseable: false,
	}

	// NOTE: See note for CardSet.Card()
	card := ps.CardSpec.Build()
	for i := 0; i < pileSize; i++ {
		pile.AddCard(card.Clone())
	}

	return pile
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

func (s PileSizeSpec) Build(numPlayers int) int {
	size := s.DefaultPileSize

	// NOTE: Should all FooSpec structs expost a Build() method?
	// In other words, if something is being referred to as a FooSpec,
	// should it not be providing the means to build a Foo?
	// If that's the case, this should actually be something like:
	//
	// pSize := p.Build(numPlayers)
	// if pSize > -1 {
	//   size = pSize
	// }
	//
	// Rigidly enforcing this would mean parent specs need not be required
	// to understand the structure of the specs they compose. It would make
	// other things a little awkward though.
	for _, p := range s.PlayerCountPileSizeSpecs {
		if p.PlayerCount == numPlayers {
			size = p.PileSize
		}
	}

	return size
}

// PlayerCountPileSizeSpec describes how the size of a pile relates
// to player count.
type PlayerCountPileSizeSpec struct {
	PlayerCount int
	PileSize    int
}

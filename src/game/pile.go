package game

import (
	"math/rand"
	"time"
)

// A Pile is an ordered collection of Cards.
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
//
// NOTE: Specifications should be considered part of the
// domain and therefore not include struct tags related
// to yaml parsing.
//
// TODO: Following on the above, move yaml parsing somewhere else.
type PileSpec struct {
	CardSpec CardSpec `yaml:"card"`

	PileSizeSpec PileSizeSpec `yaml:"size"`
}

func (ps PileSpec) Build(numPlayers int) Pile {
	pileSize := ps.PileSizeSpec.Build(numPlayers)

	pile := Pile{
		Countable:  true,
		Faceup:     true,
		Browseable: false,
	}

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
	DefaultPileSize int `yaml:"default"`

	// PlayerCountPileSizeSpecs describes pile size specifications
	// based on player count.
	//
	// TODO: Revisit the naming for this...
	PlayerCountPileSizeSpecs []PlayerCountPileSizeSpec `yaml:"playerCount"`
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
	PlayerCount int `yaml:"players"`
	PileSize    int `yaml:"pileSize"`
}

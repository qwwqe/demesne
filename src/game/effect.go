package game

type Effect struct {
	GainAction   *EffectGainAction
	GainTreasure *EffectGainTreasure
	GainCard     *EffectGainCard
	Discard      *EffectDiscard
	Draw         *EffectDraw
	Trash        *EffectTrash
	Ignore       *EffectIgnore
	Play         *EffectPlay
	Attack       *EffectAttack
	Reveal       *EffectReveal

	Optional *Effect
	Sequence []Effect
}

type EffectType string

type EffectLocation struct {
	Identifier EffectLocationIdentifier
	Specifier  EffectLocationSpecifier
}

type EffectLocationIdentifier string

const (
	EffectLocationIdentifierHand    EffectLocationIdentifier = "hand"
	EffectLocationIdentifierDiscard EffectLocationIdentifier = "discard"
	EffectLocationIdentifierDeck    EffectLocationIdentifier = "deck"
)

type EffectLocationSpecifier string

const (
	EffectLocationSpecifierTop    EffectLocationSpecifier = "top"
	EffectLocationSpecifierBottom EffectLocationSpecifier = "bottom"
	EffectLocationSpecifierAny    EffectLocationSpecifier = "any"
)

func NewEffectLocation(identifier EffectLocationIdentifier, specifier EffectLocationSpecifier) *EffectLocation {
	return &EffectLocation{
		Identifier: identifier,
		Specifier:  specifier,
	}
}

func EffectLocationDeckTop() *EffectLocation {
	return NewEffectLocation(
		EffectLocationIdentifierDeck,
		EffectLocationSpecifierTop,
	)
}

func EffectLocationDiscardAny() *EffectLocation {
	return NewEffectLocation(
		EffectLocationIdentifierDiscard,
		EffectLocationSpecifierAny,
	)
}

func EffectLocationDiscardTop() *EffectLocation {
	return NewEffectLocation(
		EffectLocationIdentifierDiscard,
		EffectLocationSpecifierTop,
	)
}

func EffectLocationHandAny() *EffectLocation {
	return NewEffectLocation(
		EffectLocationIdentifierHand,
		EffectLocationSpecifierAny,
	)
}

type EffectGainAction struct {
	Amount AmountFixed
}

const EffectTypeGainAction EffectType = "gainAction"

type EffectGainTreasure struct {
	Amount AmountFixed
}

const EffectTypeGainTreasure EffectType = "gainTreasure"

type EffectGainCard struct {
	Cost *Amount
	Name *string
}

const EffectTypeGainCard EffectType = "gainCard"

type EffectDiscard struct {
	Amount Amount
}

const EffectTypeDiscard EffectType = "discard"

type EffectDraw struct {
	Amount Amount
}

const EffectTypeDraw EffectType = "draw"

type EffectTrash struct {
	From   EffectLocation
	Amount Amount
}

const EffectTypeTrash EffectType = "trash"

type EffectIgnore struct {
}

type EffectPlay struct {
	Type *CardType
}

type EffectAttack struct{}

type EffectReveal struct{}

type EffectSpec Effect

func (s EffectSpec) Build() Effect {
	return Effect(s)
}

type ReactionTargetRole string

const (
	ReactionTargetSelf  ReactionTargetRole = "self"
	ReactionTargetOther ReactionTargetRole = "other"
)

type Reaction struct {
	TargetRole ReactionTargetRole
	// TODO: this should be an Event, not an Effect
	TargetEffect    Effect
	ReactionEffects []Effect
}

type ReactionSpec Reaction

func (s ReactionSpec) Build() Reaction {
	return Reaction(s)
}

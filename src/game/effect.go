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
	View         *EffectView
	Take         *EffectTake

	Reaction  *Reaction
	Condition *EffectCondition
	Optional  []Effect
	Sequence  []Effect
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

type EffectDiscard EffectTake

const EffectTypeDiscard EffectType = "discard"

type EffectDraw EffectTake

func EffectStandardDraw(amount AmountFixed) *EffectDraw {
	return &EffectDraw{
		From:   *EffectLocationDeckTop(),
		To:     *EffectLocationHandAny(),
		Amount: *BasicAmount(amount),
	}
}

const EffectTypeDraw EffectType = "draw"

type EffectTrash struct {
	From   EffectLocation
	Amount Amount
}

const EffectTypeTrash EffectType = "trash"

type EffectIgnore struct{}

const EffectTypeIgnore EffectType = "ignore"

type EffectPlay struct {
	Types []CardType
	Names []string
}

const EffectTypePlay EffectType = "play"

type EffectAttack struct{}

const EffectTypeAttack EffectType = "attack"

type EffectReveal struct {
	Types []CardType
	This  *bool
}

const EffectTypeReveal EffectType = "reveal"

type EffectView struct {
	Target EffectLocation
}

const EffectTypeView EffectType = "view"

type EffectTake struct {
	Amount Amount
	From   EffectLocation
	To     EffectLocation
}

const EffectTypeTake EffectType = "take"

// TODO: Merge "filter" or "query" structures (such as the one for Play)
// together into a general purpose matching formulation
type EffectCondition struct {
}

const EffectTypeCondition EffectType = "condition"

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
	Target  ReactionTarget
	Effects []Effect
}

type ReactionTarget struct {
	Role ReactionTargetRole
	// TODO: This should be an Event
	Effect Effect
	Limit  uint
}

type ReactionSpec Reaction

func (s ReactionSpec) Build() Reaction {
	return Reaction(s)
}

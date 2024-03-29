package game

type Effect struct {
	GainAction   *EffectGainAction
	GainBuy      *EffectGainBuy
	GainTreasure *EffectGainTreasure
	GainCard     *EffectGainCard
	GainVictory  *EffectGainVictory
	Discard      *EffectDiscard
	Draw         *EffectDraw
	Trash        *EffectTrash
	Ignore       *EffectIgnore
	Play         *EffectPlay
	Attack       *EffectAttack
	Interaction  *EffectInteraction
	Reveal       *EffectReveal
	View         *EffectView
	Take         *EffectTake
	SetAside     *EffectSetAside

	Reaction      *Reaction
	CardCondition *EffectCardCondition
	Optional      []Effect
	Sequence      []Effect
}

type EffectType string

type EffectLocation struct {
	Identifier EffectLocationIdentifier
	Specifier  EffectLocationSpecifier
}

type EffectLocationIdentifier string

const (
	EffectLocationIdentifierHand       EffectLocationIdentifier = "hand"
	EffectLocationIdentifierDiscard    EffectLocationIdentifier = "discard"
	EffectLocationIdentifierDeck       EffectLocationIdentifier = "deck"
	EffectLocationIdentifierPossession EffectLocationIdentifier = "possession"
	EffectLocationIdentifierSupply     EffectLocationIdentifier = "supply"
	EffectLocationIdentifierAside      EffectLocationIdentifier = "aside"
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

func EffectLocationSupplyTop() *EffectLocation {
	return NewEffectLocation(
		EffectLocationIdentifierSupply,
		EffectLocationSpecifierTop,
	)
}

func EffectLocationAsideAny() *EffectLocation {
	return NewEffectLocation(
		EffectLocationIdentifierAside,
		EffectLocationSpecifierAny,
	)
}

type EffectResult struct {
	Effect *EffectType
	This   *bool
}

type EffectCardCost struct {
	Treasure *Amount
}

type EffectGainAction struct {
	Amount AmountFixed
}

const EffectTypeGainAction EffectType = "gainAction"

type EffectGainBuy struct {
	Amount AmountFixed
}

const EffectTypeGainBuy EffectType = "gainBuy"

type EffectGainVictory struct {
	Amount Amount
}

const EffectTypeGainVictory EffectType = "gainVictory"

type EffectGainTreasure struct {
	Amount AmountFixed
}

const EffectTypeGainTreasure EffectType = "gainTreasure"

type EffectGainCard EffectTake

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

type EffectTrash EffectTake

const EffectTypeTrash EffectType = "trash"

type EffectIgnore struct{}

const EffectTypeIgnore EffectType = "ignore"

type EffectPlay struct {
	From   *EffectLocation
	Types  []CardType
	Names  []string
	Result *EffectResult
	Times  *Amount
}

const EffectTypePlay EffectType = "play"

type EffectAttackTarget struct {
	Other *bool
}

func EffectAttackTargetOther() *EffectAttackTarget {
	t := true
	return &EffectAttackTarget{&t}
}

type EffectAttack struct {
	Target  EffectAttackTarget
	Effects []Effect
}

const EffectTypeAttack EffectType = "attack"

type EffectInteraction EffectAttack

const EffectTypeInteraction EffectType = "interaction"

type EffectReveal struct {
	Types  []CardType
	This   *bool
	Hand   *bool
	Amount Amount
	From   *EffectLocation
}

const EffectTypeReveal EffectType = "reveal"

type EffectView struct {
	Target EffectLocation
}

const EffectTypeView EffectType = "view"

type EffectSetAside EffectTake

const EffectTypeSetAside EffectType = "setAside"

type EffectTake struct {
	From     EffectLocation
	To       EffectLocation
	Amount   Amount
	Target   *EffectCardConditionTarget
	Criteria *EffectCardConditionCriteria
	PerCard  *Effect // TODO: Name this better.
}

const EffectTypeTake EffectType = "take"

// TODO: Merge "filter" or "query" structures (such as the one for Play, Gain, etc)
// together into a general purpose matching formulation
type EffectCardCondition struct {
	Target   EffectCardConditionTarget
	Criteria EffectCardConditionCriteria
	Effects  []Effect
}

type EffectCardConditionTarget struct {
	Result *EffectResult
}

type EffectCardConditionCriteria struct {
	Types  []CardType
	Names  []string
	Cost   *EffectCardCost
	Amount *Amount
	Not    *EffectCardConditionCriteria
	Result *EffectResult
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

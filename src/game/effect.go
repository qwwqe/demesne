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
}

type EffectType string

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
	Amount Amount
}

const EffectTypeTrash EffectType = "trash"

type EffectIgnore struct{}

type EffectPlay struct{}

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
	TargetRole     ReactionTargetRole
	TargetEffect   Effect
	ReactionEffect Effect
}

type ReactionSpec Reaction

func (s ReactionSpec) Build() Reaction {
	return Reaction(s)
}

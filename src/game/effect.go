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

type EffectGainAction struct {
	Amount AmountFixed
}

type EffectGainTreasure struct {
	Amount AmountFixed
}

type EffectGainCard struct {
	Cost *Amount
	Name *string
}

type EffectDiscard struct{}

type EffectDraw struct{}

type EffectTrash struct{}

type EffectIgnore struct{}

type EffectPlay struct{}

type EffectAttack struct{}

type EffectReveal struct{}

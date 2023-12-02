package game

type Effect struct {
	GainAction   *EffectGainAction
	GainTreasure *EffectGainTreasure
	GainCard     *EffectGainCard
	GainVictory  *EffectGainVictory
	Discard      *EffectDiscard
	Draw         *EffectDraw
	Trash        *EffectTrash
	Ignore       *EffectIgnore
	Play         *EffectPlay
	Attack       *EffectAttack
	Reveal       *EffectReveal
}

type EffectGainAction struct {
	Amount Amount
}

type EffectGainTreasure struct {
	Amount Amount
}

type EffectGainCard struct {
	Cost *Amount
	Name *string
}

type EffectGainVictory struct{}

type EffectDiscard struct{}

type EffectDraw struct{}

type EffectTrash struct{}

type EffectIgnore struct{}

type EffectPlay struct{}

type EffectAttack struct{}

type EffectReveal struct{}

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
	Amount Amount
}

type EffectGainVictory struct {
	Amount Amount
}

type EffectDiscard struct {
	Amount Amount
}

type EffectDraw struct {
	Amount Amount
}

type EffectTrash struct{}

type EffectIgnore struct{}

type EffectPlay struct{}

type EffectAttack struct{}

type EffectReveal struct{}

package game

import (
	"testing"
)

func IntPtr(i int) *int {
	return &i
}

func TestCardEffects(t *testing.T) {
	cellar := CardSpec{
		Name:  "cellar",
		Cost:  CostSpec{Treasure: 2},
		Types: []CardTypeSpec{CardTypeSpec(CardTypeAction)},
		Effects: []EffectSpec{
			EffectSpec{GainAction: &EffectGainAction{1}},
			EffectSpec{
				Discard: &EffectDiscard{
					Amount: Amount{
						Range: &AmountRange{Min: IntPtr(0)},
					},
				},
			},
			EffectSpec{
				Draw: &EffectDraw{
					Amount: Amount{
						Result: &AmountResult{EffectTypeDiscard},
					},
				},
			},
		},
	}

	cellar.Build()
}

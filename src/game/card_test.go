package game

import (
	"testing"
)

func IntPtr(i int) *int {
	return &i
}

func TestCardEffects(t *testing.T) {
	t.Run("cellar", func(t *testing.T) {
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
	})

	t.Run("chapel", func(t *testing.T) {
		chapel := CardSpec{
			Name:  "chapel",
			Cost:  CostSpec{Treasure: 2},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeAction)},
			Effects: []EffectSpec{
				EffectSpec{
					Trash: &EffectTrash{
						Amount{
							Range: &AmountRange{
								Max: IntPtr(4),
							},
						},
					},
				},
			},
		}

		chapel.Build()
	})

	t.Run("moat", func(t *testing.T) {
		moat := CardSpec{
			Name: "moat",
			Cost: CostSpec{Treasure: 2},
			Types: []CardTypeSpec{
				CardTypeSpec(CardTypeAction),
				CardTypeSpec(CardTypeReaction),
			},
			Effects: []EffectSpec{
				EffectSpec{},
			},
		}

		moat.Build()
	})

	t.Run("harbinger", func(t *testing.T) {
		harbinger := CardSpec{}

		harbinger.Build()
	})

	t.Run("merchant", func(t *testing.T) {
		merchant := CardSpec{}

		merchant.Build()
	})
}

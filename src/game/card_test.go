package game

import (
	"testing"
)

func IntPtr(i int) *int {
	return &i
}

func UintPtr(u uint) *uint {
	return &u
}

func StrPtr(s string) *string {
	return &s
}

func TestCardEffects(t *testing.T) {
	t.Run("cellar", func(t *testing.T) {
		cellar := CardSpec{
			Name:  "cellar",
			Cost:  CostSpec{Treasure: 2},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeAction)},
			ActionEffects: []EffectSpec{
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
			ActionEffects: []EffectSpec{
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
			ActionEffects: []EffectSpec{
				{
					Draw: &EffectDraw{
						Amount{
							Fixed: (*AmountFixed)(UintPtr(2)),
						},
					},
				},
			},
			ReactionEffects: []ReactionSpec{
				ReactionSpec{
					TargetRole: ReactionTargetOther,
					TargetEffect: Effect{
						Play: &EffectPlay{
							Type: (*CardType)(StrPtr(string(CardTypeAttack))),
						},
					},
					ReactionEffects: []Effect{
						{
							Optional: &Effect{
								Sequence: []Effect{
									{Ignore: &EffectIgnore{}},
								},
							},
						},
					},
				},
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

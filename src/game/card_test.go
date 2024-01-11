package game

import (
	"testing"
)

func Ptr[T any](v T) *T {
	return &v
}

func TestCardEffects(t *testing.T) {
	t.Run("cellar", func(t *testing.T) {
		cellar := CardSpec{
			Name:  "cellar",
			Cost:  CostSpec{Treasure: 2},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeAction)},
			ActionEffects: []EffectSpec{
				{GainAction: &EffectGainAction{1}},
				{
					Discard: &EffectDiscard{
						Amount: Amount{
							Range: &AmountRange{Min: Ptr(0)},
						},
					},
				},
				{
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
				{
					Trash: &EffectTrash{
						Amount{
							Range: &AmountRange{
								Max: Ptr(4),
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
							Fixed: Ptr(AmountFixed(2)),
						},
					},
				},
			},
			ReactionEffects: []ReactionSpec{
				{
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

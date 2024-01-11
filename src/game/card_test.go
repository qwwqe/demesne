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
						From: *EffectLocationHandAny(),
						Amount: Amount{
							Range: &AmountRange{Min: Ptr(0)},
						},
					},
				},
				{
					Draw: &EffectDraw{
						From: *EffectLocationDeckTop(),
						To:   *EffectLocationHandAny(),
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
						From: *EffectLocationHandAny(),
						Amount: Amount{
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
					Draw: EffectStandardDraw(2),
				},
			},
			ReactionEffects: []ReactionSpec{
				{
					Target: ReactionTarget{
						Role: ReactionTargetOther,
						Effect: Effect{
							Play: &EffectPlay{
								Types: []CardType{CardTypeAttack},
							},
						},
					},
					Effects: []Effect{
						{
							Optional: []Effect{
								{Reveal: &EffectReveal{This: Ptr(true)}},
								{Ignore: &EffectIgnore{}},
							},
						},
					},
				},
			},
		}

		moat.Build()
	})

	t.Run("harbinger", func(t *testing.T) {
		harbinger := CardSpec{
			Name:  "harbinger",
			Cost:  CostSpec{Treasure: 3},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeAction)},
			ActionEffects: []EffectSpec{
				{
					Draw: EffectStandardDraw(1),
				},
				{
					GainAction: &EffectGainAction{1},
				},
				{
					View: &EffectView{
						Target: *EffectLocationDiscardAny(),
					},
				},
				{
					Optional: []Effect{{Take: &EffectTake{
						From: *EffectLocationDiscardAny(),
						To:   *EffectLocationDeckTop(),
						Amount: Amount{
							Fixed: Ptr(AmountFixed(1)),
						},
					}}},
				},
			},
		}

		harbinger.Build()
	})

	t.Run("merchant", func(t *testing.T) {
		merchant := CardSpec{
			Name:  "merchant",
			Cost:  CostSpec{Treasure: 3},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeAction)},
			ActionEffects: []EffectSpec{
				{
					Draw: EffectStandardDraw(1),
				},
				{
					GainAction: &EffectGainAction{1},
				},
				{
					Reaction: &Reaction{
						Target: ReactionTarget{
							Role: ReactionTargetSelf,
							Effect: Effect{
								Play: &EffectPlay{
									Names: []string{"silver"},
								},
							},
							Limit: 1,
						},
						Effects: []Effect{{
							GainTreasure: &EffectGainTreasure{
								Amount: 1,
							},
						}},
					},
				},
			},
		}

		merchant.Build()
	})
}

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
						From:   *EffectLocationDiscardAny(),
						To:     *EffectLocationDeckTop(),
						Amount: *BasicAmount(1),
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

	t.Run("vassal", func(t *testing.T) {
		vassal := CardSpec{
			Name:  "vassal",
			Cost:  CostSpec{Treasure: 3},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeAction)},
			ActionEffects: []EffectSpec{
				{GainTreasure: &EffectGainTreasure{Amount: 2}},
				{Discard: &EffectDiscard{
					From:   *EffectLocationDeckTop(),
					To:     *EffectLocationDiscardTop(),
					Amount: *BasicAmount(1),
				}},
				{
					CardCondition: &EffectCardCondition{
						Target: EffectCardConditionTarget{
							Result: &EffectResult{
								Effect: Ptr(EffectTypeDiscard),
							},
						},
						Criteria: EffectCardConditionCriteria{
							Types: []CardType{CardTypeAction},
						},
						Effects: []Effect{{
							Optional: []Effect{{
								Play: &EffectPlay{
									Result: &EffectResult{
										Effect: Ptr(EffectTypeDiscard),
									},
								},
							}},
						}},
					},
				},
			},
		}

		vassal.Build()
	})

	t.Run("village", func(t *testing.T) {
		village := CardSpec{
			Name:  "village",
			Cost:  CostSpec{Treasure: 3},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeAction)},
			ActionEffects: []EffectSpec{
				{Draw: EffectStandardDraw(1)},
				{GainAction: &EffectGainAction{2}},
			},
		}

		village.Build()
	})

	t.Run("workshop", func(t *testing.T) {
		workshop := CardSpec{
			Name:  "workshop",
			Cost:  CostSpec{Treasure: 3},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeAction)},
			ActionEffects: []EffectSpec{
				{
					GainCard: &EffectGainCard{
						From:   *EffectLocationSupplyTop(),
						To:     *EffectLocationDiscardTop(),
						Amount: Amount{Fixed: Ptr(AmountFixed(1))},
						Criteria: &EffectCardConditionCriteria{
							Cost: &EffectCardCost{
								Treasure: &Amount{
									Range: &AmountRange{
										Max: Ptr(4),
									},
								},
							},
						},
					},
				},
			},
		}

		workshop.Build()
	})

	t.Run("bureaucrat", func(t *testing.T) {
		bureaucrat := CardSpec{
			Name:  "bureaucrat",
			Cost:  CostSpec{Treasure: 4},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeAttack), CardTypeSpec(CardTypeAction)},
			ActionEffects: []EffectSpec{
				{
					GainCard: &EffectGainCard{
						To:     *EffectLocationDeckTop(),
						Amount: Amount{Fixed: Ptr(AmountFixed(1))},
						Criteria: &EffectCardConditionCriteria{
							Names: []string{"silver"},
						},
					},
				},
				{
					Attack: &EffectAttack{
						Target: EffectAttackTarget{Other: Ptr(true)},
						Effects: []Effect{
							{
								Reveal: &EffectReveal{
									Types:  []CardType{CardTypeVictory},
									Amount: Amount{Fixed: Ptr(AmountFixed(1))},
								},
							},
							{
								Discard: &EffectDiscard{
									To: *EffectLocationDeckTop(),
									Target: &EffectCardConditionTarget{
										Result: &EffectResult{
											Effect: Ptr(EffectTypeReveal),
										},
									},
								},
							},
							{
								CardCondition: &EffectCardCondition{
									Target: EffectCardConditionTarget{
										Result: &EffectResult{
											Effect: Ptr(EffectTypeDiscard),
										},
									},
									Criteria: EffectCardConditionCriteria{
										Amount: &Amount{
											Fixed: Ptr(AmountFixed(0)),
										},
									},
									Effects: []Effect{{
										Reveal: &EffectReveal{
											Hand: Ptr(true),
										},
									}},
								},
							},
						},
					},
				},
			},
		}

		bureaucrat.Build()
	})

	t.Run("gardens", func(t *testing.T) {
		gardens := CardSpec{
			Name:  "gardens",
			Cost:  CostSpec{Treasure: 4},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeVictory)},
			ScoringEffects: []EffectSpec{{
				GainVictory: &EffectGainVictory{
					Amount: Amount{
						Relative: &AmountRelative{
							Target: AmountRelativeTarget{
								LocationIdentifier: Ptr(EffectLocationIdentifierPossession),
							},
							Divider: Ptr(10),
						},
					},
				}},
			},
		}

		gardens.Build()
	})

	t.Run("militia", func(t *testing.T) {
		militia := CardSpec{
			Name:  "militia",
			Cost:  CostSpec{Treasure: 4},
			Types: []CardTypeSpec{CardTypeSpec(CardTypeAction), CardTypeSpec(CardTypeAttack)},
			ActionEffects: []EffectSpec{
				{GainTreasure: &EffectGainTreasure{Amount: 1}},
				{Attack: &EffectAttack{}},
			},
		}

		militia.Build()
	})

}

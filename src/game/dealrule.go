package game

// DealRule describes a dealing rule.
type DealRule struct {
	PileName string

	AmountPerPlayer int

	DepletePile bool
}

func (d DealRule) Deal(p *Pile) []Card {
	if d.DepletePile {
		return p.Draw(d.AmountPerPlayer)
	}

	cards := make([]Card, d.AmountPerPlayer)
	card := p.Top()
	for i := 0; i < d.AmountPerPlayer; i++ {
		cards = append(cards, card.Clone())
	}

	return cards
}

// DealRuleSpec describes how a dealing rule is created.
type DealRuleSpec struct {
	DealRule
}

func (s DealRuleSpec) Build() DealRule {
	return s.DealRule
}

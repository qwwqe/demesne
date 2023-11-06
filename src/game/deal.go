package game

type Dealer interface {
	// Deal a single hand.
	Deal(p *Pile) []Card
}

// DestructiveDealer deals by directly drawing from a pile.
type DestructiveDealer struct {
	amount int
}

// Deal a single hand.
func (d DestructiveDealer) Deal(p *Pile) []Card {
	return p.Draw(d.amount)
}

var _ Dealer = DestructiveDealer{}

// NonDestructiveDealer deals by cloning the top card of a pile.
type NonDestructiveDealer struct {
	amount int
}

// Deal a single hand.
func (d NonDestructiveDealer) Deal(p *Pile) []Card {
	card := p.Top()
	if card == nil {
		return nil
	}

	cs := make([]Card, d.amount)
	for i := 0; i < len(cs); i++ {
		cs[i] = card.Clone()
	}

	return cs
}

var _ Dealer = NonDestructiveDealer{}

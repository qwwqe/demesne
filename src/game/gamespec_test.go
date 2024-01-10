package game

// A simple end condition based on supply pile exhaustion.
//
// This is mostly intended for reference until a more comprehensive
// framework for dynamic definition is established.
func basicSupplyEndCondition(g Game) bool {
	supplyPilesExhausted := 0

	for _, p := range g.Supply {
		if p.Size() == 0 {
			supplyPilesExhausted++
		}
	}

	if len(g.Players) < 5 {
		return supplyPilesExhausted >= 3
	}

	return supplyPilesExhausted >= 4
}

type provincePileSpec struct{}

func (ps provincePileSpec) newCard() Card {
	return Card{
		Name: "province",
	}
}

func (ps provincePileSpec) id() string {
	return "province"
}

func (ps provincePileSpec) Build(numPlayers int) Pile {
	pileSize := 12
	if numPlayers == 2 {
		pileSize = 8
	} else if numPlayers == 5 {
		pileSize = 15
	} else if numPlayers == 6 {
		pileSize = 18
	}

	pile := Pile{
		Countable:  true,
		Faceup:     true,
		Browseable: false,
	}

	// NOTE: See note for CardSet.Card()
	card := ps.newCard()
	for i := 0; i < pileSize; i++ {
		pile.AddCard(card.Clone())
	}

	return pile
}

func (cs provincePileSpec) EndConditions() []EndCondition {
	return []EndCondition{
		EmptyPileEndCondition{cs.id()},
	}
}

func (cs provincePileSpec) Deal(*Pile) []Card {
	return nil
}

type estatePileSpec struct{}

func (ps estatePileSpec) id() string {
	return "estate"
}

func (ps estatePileSpec) newCard() Card {
	return Card{
		Name: "estate",
	}
}

func (ps estatePileSpec) Build(numPlayers int) Pile {
	pileSize := 8
	if numPlayers != 2 {
		pileSize = 12
	}

	amountPerPlayer := 3
	pileSize += amountPerPlayer * numPlayers

	pile := Pile{
		Countable:  true,
		Faceup:     true,
		Browseable: false,
	}

	// NOTE: See note for CardSet.Card()
	card := ps.newCard()
	for i := 0; i < pileSize; i++ {
		pile.AddCard(card.Clone())
	}

	return pile
}

func (cs estatePileSpec) Deal(pile *Pile) []Card {
	amountPerPlayer := 3
	return pile.Draw(amountPerPlayer)
}

func (cs estatePileSpec) EndConditions() []EndCondition {
	return []EndCondition{}
}

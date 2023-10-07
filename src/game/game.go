package game

import "github.com/qwwqe/demesne/src/card"

// A game of Demesne.
type game struct {
	// Id uniquely identifies a Game.
	Id string

	// The current turn.
	//
	// NOTE: This state information is actually specific
	// to a certain stage of the game, and should probably
	// be stored within that stage, rather than within the
	// game itself.
	Turn int

	// The current stage.
	//
	// TODO: Evaluate necessity of explicit stages in a game.
	Stage Stage

	// The Players in the Game.
	//
	// The order of this array determines turn playing order.
	// In other words, the current player is given by the following relation:
	// 		currentPlayer := Players[Turn % len(Players)]
	Players []Player

	// The Trash.
	Trash card.Pile

	// The Supply.
	Supply
}

// The Supply is the collection of all card Piles which can be
// directly purchased from in a given game of Demesne.
//
// NOTE: It may be worth considering implementing the Supply
// in a way that makes determining existence of a card in the Supply
// easier and further expansion more convenient.
type Supply struct {
	BaseCards    []card.Pile
	KingdomCards []card.Pile
}

// A Stage is a distinct state in the life cycle of a game.
//
// NOTE: Stub.
// NOTE: It may not be necessary to categorize the game into stages here.
type Stage interface{}

// Game builder.
type Builder struct {
	game game
}

// Add Player to configured game.
func (b *Builder) WithPlayer(player Player) *Builder {
	b.game.Players = append(b.game.Players, player)
	return b
}

// Add Kingdom pile to configured game.
func (b *Builder) WithKingdom(kingdom card.Pile) *Builder {
	b.game.KingdomCards = append(b.game.KingdomCards, kingdom)
	return b
}

// Add Base pile to configured game.
func (b *Builder) WithBase(base card.Pile) *Builder {
	b.game.BaseCards = append(b.game.BaseCards, base)
	return b
}

// Return the configured game.
//
// TODO: Validate before returning.
func (b *Builder) Build() (game, error) {
	return b.game, nil
}

package game

import (
	"fmt"

	"github.com/qwwqe/demesne/src/card"
)

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

// Validate a game of Demesne.
//
// NOTE: Logic pertaining to game rules shouldn't really
// just be hardcoded here... maybe use a specification pattern?
// Furthermore, validation should probably also cover pile size of
// base and victory cards with respect to player count (i.e. 8 for a two-player
// game, 12 for a three-player game, and so on). It seems unreasonable
// to include all such logic in a single validation function,
// particularly when attempting to implement functionality added by
// future expansions.
//
// TODO: Per above, implement this using a specification pattern
// or something equivalent.
func (g game) Validate() error {
	e := func(err error) error {
		return fmt.Errorf("Validate: %w", err)
	}

	if g.Id == "" {
		return e(ErrMissingId)
	}

	if len(g.Players) < 2 || len(g.Players) > 6 {
		return e(ErrInvalidPlayerCount)
	}

	if len(g.KingdomCards) != 10 {
		return e(ErrInvalidKingdomCount)
	}

	if len(g.BaseCards) != 7 {
		return e(ErrInvalidBaseCount)
	}

	return nil
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

// Return the configured game, ready to play.
//
// TODO: Deal decks, initialize hands, set the turn counter, etc.
func (b *Builder) Build() (*game, error) {
	if err := b.game.Validate(); err != nil {
		return nil, err
	}

	return &b.game, nil
}

package game

// A Player represents a participant in a game of Demesne.
// It does not attempt to implement any functionality related to the identity of
// or communication between the actual humans or programs controlling such participation.
// In other words, notions such as names, user registration, and the like are strictly
// outside the domain served by this type.
//
// TODO: Determine how Stage-specific state, such as progress in a given turn,
// might be represented here (if at all -- maybe it belongs in the Game state)
type Player struct {
	// Id uniquely identifies a Player.
	Id string

	// The Player's current hand.
	Hand []Card

	// The Player's current discard pile.
	Discard Pile

	// The Player's deck.
	Deck Pile

	// What the Player currently has in play.
	PlayArea []Card
}

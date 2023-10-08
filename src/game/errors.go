package game

import "errors"

var (
	ErrMissingId           = errors.New("missing Id")
	ErrInvalidPlayerCount  = errors.New("invalid player count")
	ErrInvalidKingdomCount = errors.New("invalid kingdom count")
	ErrInvalidBaseCount    = errors.New("invalid base count")
)

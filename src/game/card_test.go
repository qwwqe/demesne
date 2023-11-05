package game

import (
	"testing"
)

func TestMergePiles(t *testing.T) {
	deck := Pile{
		Cards:      []Card{{Id: "1"}, {Id: "2"}},
		Faceup:     false,
		Countable:  true,
		Browseable: false,
	}

	discard := Pile{
		Cards:      []Card{{Id: "3"}, {Id: "4"}},
		Faceup:     true,
		Countable:  false,
		Browseable: false,
	}

	merged := deck.Merged(discard)
}

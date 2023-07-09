package main

import (
	"errors"
	"fmt"
)

type CardKind string

const (
	CardKindAction   CardKind = "action"
	CardKindTreasure CardKind = "treasure"
	CardKindVictory  CardKind = "victory"
	CardKindCurse    CardKind = "curse"
	CardKindReaction CardKind = "reaction"
)

type CardSet string

const (
	CardSetBasic    CardSet = "basic"
	CardSetDominion CardSet = "dominion"
)

type Card struct {
	Version string     `yaml:"demesne"`
	Name    string     `yaml:"name"`
	Cost    int        `yaml:"cost"`
	Kinds   []CardKind `yaml:"kinds"`
	Set     CardSet    `yaml:"set"`
	Effects []Effect   `yaml:"effects"`
}

type EffectKind string

const (
	EffectKindAction  EffectKind = "action"
	EffectKindDiscard EffectKind = "discard"
	EffectKindDraw    EffectKind = "draw"
	EffectKindTrash   EffectKind = "trash"
)

type Effect interface {
	Kind() EffectKind
}

// type Effect struct {
// 	Action EffectAction `yaml:"action"`
// 	Kind   EffectKind
// }

type ActionEffect struct {
	Amount int `yaml:"amount"`
}

func (ae ActionEffect) Kind() EffectKind {
	return EffectKindAction
}

type EffectAction struct {
	Amount int `yaml:"amount"`
}

const (
	versionCardFieldName string = "demesne"
	nameCardFieldName           = "name"
	costCardFieldName           = "cost"
	kindsCardFieldName          = "types"
	setCardFieldName            = "set"
	effectsCardFieldName        = "effects"
)

func cardFieldKindError(field, expectedKind string) error {
	return errors.New(fmt.Sprintf(`"%s" should be %s`, field, expectedKind))
}

func NewCardFromMap(m map[any]any) (*Card, error) {
	var c Card

	if version, ok := m[versionCardFieldName]; ok {
		s, ok := version.(string)
		if !ok {
			return nil, cardFieldKindError(versionCardFieldName, "string")
		}

		c.Version = s
	}

	if name, ok := m[nameCardFieldName]; ok {
		s, ok := name.(string)
		if !ok {
			return nil, cardFieldKindError(nameCardFieldName, "string")
		}

		c.Name = s
	}

	if cost, ok := m[costCardFieldName]; ok {
		i, ok := cost.(int)
		if !ok {
			return nil, cardFieldKindError(costCardFieldName, "int")
		}

		c.Cost = i
	}

	if set, ok := m[setCardFieldName]; ok {
		s, ok := set.(string)
		if !ok {
			return nil, cardFieldKindError(setCardFieldName, "string")
		}

		c.Set = CardSet(s)
	}

	if kinds, ok := m[kindsCardFieldName]; ok {
		ts, ok := kinds.([]interface{})
		if !ok {
			return nil, cardFieldKindError(kindsCardFieldName, "[]string")
		}

		for _, t := range ts {
			cardKind, ok := t.(string)
			if !ok {
				return nil, cardFieldKindError(kindsCardFieldName, "[]string")
			}
			c.Kinds = append(c.Kinds, CardKind(cardKind))
		}
	}

	return &c, nil
}

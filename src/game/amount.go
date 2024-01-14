package game

type AmountFixed uint

type AmountRange struct {
	Min *int
	Max *int
}

type AmountAll bool

type AmountRelativeTarget struct {
	Card               *TargetCard
	LocationIdentifier *EffectLocationIdentifier
}

type AmountRelative struct {
	Target AmountRelativeTarget
	Range  *AmountRange

	// TODO: These names are not intuitive nor reflective of what
	// the card text usually suggests
	Multiplier *int
	Divider    *int
}

type Amount struct {
	Fixed    *AmountFixed
	Range    *AmountRange
	All      *AmountAll
	Relative *AmountRelative
	Result   *AmountResult
}

func BasicAmount(amount AmountFixed) *Amount {
	return &Amount{
		Fixed: &amount,
	}
}

type TargetCard struct {
	Name *string
	Type *CardType
}

// So dirty.
type AmountResult struct {
	Effect EffectType
}

package game

type AmountFixed uint

type AmountRange struct {
	Min *int
	Max *int
}

type AmountAll bool

type AmountRelativeTarget struct {
	Result             *EffectResult
	Card               *TargetCard
	LocationIdentifier *EffectLocationIdentifier
}

type AmountRelativeUnit string

const (
	AmountRelativeUnitCard AmountRelativeUnit = "card"
	AmountRelativeUnitPile AmountRelativeUnit = "pile"
)

type AmountRelative struct {
	Target AmountRelativeTarget
	Unit   AmountRelativeUnit

	Range     *AmountRange
	Condition *EffectCardConditionCriteria

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
	Until    *AmountUntil
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

type AmountUntil struct {
	LocationIdentifier EffectLocationIdentifier
	Amount             Amount
}

// So dirty.
type AmountResult struct {
	Effect EffectType
}

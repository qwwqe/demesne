package game

type AmountFixed uint

type AmountRange struct {
	Min uint
	Max uint
}

type AmountAll bool

type Amount struct {
	Fixed *AmountFixed
	Range *AmountRange
	All   *AmountAll
}

// type AnyAmount

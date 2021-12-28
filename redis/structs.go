package redis

type Z struct {
	Score  float64
	Member interface{}
}

type ZSlice []Z

type ZStore struct {
	Keys    []string
	Weights []float64
	// Can be SUM, MIN or MAX.
	Aggregate string
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

type BitCountArgs struct {
	Start, End int64
}

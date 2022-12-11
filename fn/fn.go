package fn

func Map[X, Y any](in []X, mapper func(X) Y) []Y {
	out := make([]Y, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = mapper(in[i])
	}
	return out
}

func Filter[T any](in []T, predicate func(T) bool) []T {
	out := make([]T, 0)
	for _, x := range in {
		if predicate(x) {
			out = append(out, x)
		}
	}
	return out
}

func Reduce[T any](in []T, f func(acc, x T) T, initial T) T {
	out := initial
	for _, x := range in {
		out = f(out, x)
	}
	return out
}

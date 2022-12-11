// Package fn provides some useful helper functions.
package fn

// Map returns a slice consisting of the results of applying the given mapper
// function to the elements of the input slice.
func Map[X, Y any](in []X, mapper func(X) Y) []Y {
	out := make([]Y, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = mapper(in[i])
	}
	return out
}

// Filter return a slice consisting of the elements of the input slice that
// match the given predicate.
func Filter[T any](in []T, predicate func(T) bool) []T {
	out := make([]T, 0)
	for _, x := range in {
		if predicate(x) {
			out = append(out, x)
		}
	}
	return out
}

// Reduce performs a reduction on the elements of the slice using the provided
// associative accumlation function and identity value.
func Reduce[T any](in []T, f func(acc, x T) T, identity T) T {
	out := identity
	for _, x := range in {
		out = f(out, x)
	}
	return out
}

// Any returns true if any elements of the input slice match the provided
// predicate.
func Any[T any](in []T, predicate func(T) bool) bool {
	for _, x := range in {
		if predicate(x) {
			return true
		}
	}
	return false
}

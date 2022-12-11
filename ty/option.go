package ty

import "encoding/json"

// Optional is the interface that wraps methods for optional values.
type Optional[T any] interface {
	// Get returns the value and a boolean to indicate if the value is present.
	Get() (T, bool)

	// OrElse returns the value if it is present; otherwise, it returns the
	// provided argument.
	OrElse(T) T

	// IsPresent indicates if the value is present.
	IsPresent() bool

	// Unwrap returns the value if it is present; otherwise, it returns the zero
	// value.
	Unwrap() T
}

// An Option represents an optional value of type T and implements [ty.Optional]
// interface.
type Option[T any] struct {
	value     T
	isPresent bool
}

// None returns an [ty.Option] that represents the non-existent value of type T.
func None[T any]() *Option[T] {
	return &Option[T]{isPresent: false}
}

// Some returns an [ty.Option] that represent the value of type T.
func Some[T any](x T) *Option[T] {
	option := new(Option[T])
	option.isPresent = true
	option.value = x
	return option
}

// Get implements [ty.Optional.Get].
func (opt *Option[T]) Get() (x T, ok bool) {
	if !opt.IsPresent() {
		return
	}

	x, ok = opt.value, opt.isPresent
	return
}

// OrElse implements [ty.Optional.OrElse].
func (opt *Option[T]) OrElse(x T) T {
	if !opt.IsPresent() {
		return x
	}
	return opt.value
}

// IsPresent implements [ty.Optional.IsPresent].
func (opt *Option[T]) IsPresent() bool {
	return opt != nil && opt.isPresent
}

// Unwrap implements [ty.Optional.Unwrap].
func (opt *Option[T]) Unwrap() (x T) {
	if !opt.isPresent {
		return
	}

	return opt.value
}

// MarshalJSON implements [encoding/json.Marshaler].
func (opt Option[T]) MarshalJSON() ([]byte, error) {
	if !opt.IsPresent() {
		return json.Marshal(nil)
	}

	return json.Marshal(opt.value)
}

// UnmarshalJSON implements [encoding/json.Unmarshaler].
func (opt *Option[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	opt.isPresent = true
	return json.Unmarshal(data, &opt.value)
}

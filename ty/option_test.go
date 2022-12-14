package ty_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.inflx.dev/bronze/ty"
)

func TestOption(t *testing.T) {
	t.Run("Some[int]", func(t *testing.T) {
		opt := ty.Some(1)
		assert.True(t, opt.IsPresent())

		x, ok := opt.Get()
		assert.True(t, ok)
		assert.Equal(t, 1, x)

		assert.Equal(t, 1, opt.OrElse(2))
		assert.Equal(t, 1, opt.MustGet())
		assert.Equal(t, 1, opt.Unwrap())
	})

	t.Run("Some[string]", func(t *testing.T) {
		opt := ty.Some("hello")
		assert.True(t, opt.IsPresent())

		x, ok := opt.Get()
		assert.True(t, ok)
		assert.Equal(t, "hello", x)

		assert.Equal(t, "hello", opt.OrElse("hi"))
		assert.Equal(t, "hello", opt.MustGet())
		assert.Equal(t, "hello", opt.Unwrap())
	})

	t.Run("None[int]", func(t *testing.T) {
		opt := ty.None[int]()
		assert.False(t, opt.IsPresent())

		_, ok := opt.Get()
		assert.False(t, ok)

		assert.Equal(t, 2, opt.OrElse(2))
		assert.Panics(t, func() { opt.MustGet() })
		assert.Equal(t, 0, opt.Unwrap())
	})

	t.Run("None[string]", func(t *testing.T) {
		opt := ty.None[string]()
		assert.False(t, opt.IsPresent())

		_, ok := opt.Get()
		assert.False(t, ok)

		assert.Equal(t, "hi", opt.OrElse("hi"))
		assert.Panics(t, func() { opt.MustGet() })
		assert.Equal(t, "", opt.Unwrap())
	})

	t.Run("MarshalJSON Some[string]", func(t *testing.T) {
		opt := ty.Some("helloworld")
		bytes, err := json.Marshal(opt)
		assert.NoError(t, err)
		assert.Equal(t, `"helloworld"`, string(bytes))
	})

	t.Run("MarshalJSON None[string]", func(t *testing.T) {
		opt := ty.None[string]()
		bytes, err := json.Marshal(opt)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(bytes))
	})

	t.Run("UnmarshalJSON Some[string]", func(t *testing.T) {
		var opt *ty.Option[string]
		err := json.Unmarshal([]byte(`"helloworld"`), &opt)
		assert.NoError(t, err)
		assert.Equal(t, "helloworld", opt.Unwrap())
	})

	t.Run("UnmarshalJSON None[string]", func(t *testing.T) {
		var opt *ty.Option[string]
		err := json.Unmarshal([]byte(`null`), &opt)
		assert.NoError(t, err)
		assert.False(t, opt.IsPresent())
	})
}

func ExampleSome() {
	opt := ty.Some("hello")
	x, ok := opt.Get()
	fmt.Println(x, ok)
	// Output: hello true
}

func ExampleNone() {
	opt := ty.None[int]()
	x, ok := opt.Get()
	fmt.Println(x, ok)
	// Output: 0 false
}

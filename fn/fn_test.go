package fn_test

import (
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"inflx.dev/bronze/fn"
)

func TestMap(t *testing.T) {
	t.Run("map int to string", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		out := fn.Map(in, func(x int) int { return x * x })
		assert.Equal(t, []int{1, 4, 9, 16}, out)
	})

	t.Run("map float to its square", func(t *testing.T) {
		in := []float64{1, 2, 3, 4}
		out := fn.Map(in, func(x float64) float64 { return math.Pow(x, 2) })
		assert.Equal(t, []float64{1, 4, 9, 16}, out)
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter out odd numbers", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6}
		out := fn.Filter(in, func(x int) bool { return x%2 == 0 })
		assert.Equal(t, []int{2, 4, 6}, out)
	})

	t.Run("filter out nil", func(t *testing.T) {
		err1, err2 := errors.New("err1"), errors.New("err2")
		in := []error{err1, nil, err2}
		out := fn.Filter(in, func(x error) bool { return x != nil })
		assert.Equal(t, []error{err1, err2}, out)
	})
}

func TestReduce(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6}
		out := fn.Reduce(in, func(acc, x int) int { return acc + x }, 0)
		assert.Equal(t, 21, out)
	})

	t.Run("concat strings", func(t *testing.T) {
		in := []string{"Hello", "World", "!"}
		out := fn.Reduce(in, func(acc, x string) string { return acc + x }, "")
		assert.Equal(t, "HelloWorld!", out)
	})
}

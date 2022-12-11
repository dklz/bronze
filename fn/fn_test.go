package fn_test

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.inflx.dev/bronze/fn"
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

	t.Run("map nil slice", func(t *testing.T) {
		out := fn.Map(nil, func(x float64) float64 { return x })
		assert.Equal(t, []float64{}, out)
	})
}

func ExampleMap() {
	nums := []int{1, 2, 3}
	tenthPowers := fn.Map(nums, func(x int) float64 { return math.Pow(float64(x), 10) })
	fmt.Println(tenthPowers)
	// Output: [1 1024 59049]
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

	t.Run("filter nil slice", func(t *testing.T) {
		out := fn.Filter(nil, func(x int) bool { return true })
		assert.Equal(t, []int{}, out)
	})
}

func ExampleFilter() {
	nums := []int{1, 2, 3, 4, 5}
	odds := fn.Filter(nums, func(x int) bool { return x%2 != 0 })
	fmt.Println(odds)
	// Output: [1 3 5]
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

	t.Run("reduce nil slice", func(t *testing.T) {
		out := fn.Reduce(nil, func(acc, x int) int { return acc + x }, 0)
		assert.Equal(t, 0, out)
	})
}

func ExampleReduce() {
	nums := []int{1, 2, 3}
	sum := fn.Reduce(nums, func(acc, x int) int { return acc + x }, 0)
	fmt.Println(sum)
	// Output: 6
}

func TestAny(t *testing.T) {
	t.Run("string slice", func(t *testing.T) {
		in := []string{"hello", "world", "!"}
		assert.True(t, fn.Any(in, func(x string) bool { return x == "hello" }))
		assert.False(t, fn.Any(in, func(x string) bool { return len(x) > 99 }))
	})

	t.Run("int slice", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		assert.True(t, fn.Any(in, func(x int) bool { return x%2 == 0 }))
		assert.False(t, fn.Any(in, func(x int) bool { return x%5 == 0 }))
	})

	t.Run("nil slice", func(t *testing.T) {
		var in []int
		assert.False(t, fn.Any(in, func(x int) bool { return true }))
	})
}

func ExampleAny() {
	nums := []int{1, 2, 3}
	hasEvenNumber := fn.Any(nums, func(x int) bool { return x%2 == 0 })
	fmt.Println(hasEvenNumber)
	// Output: true
}

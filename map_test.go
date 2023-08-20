package gohelpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapFuncDouble(t *testing.T) {
	as := assert.New(t)
	arr := []int{1, 2, 3}
	f := func(x int) int { return x * 2 }
	result := MapFunc(arr, f)
	as.Equal([]int{2, 4, 6}, result)
}

func TestMapFuncToString(t *testing.T) {
	as := assert.New(t)
	arr := []int{1, 2, 3}
	f := func(x int) string { return fmt.Sprint(x) }
	result := MapFunc(arr, f)
	as.Equal([]string{"1", "2", "3"}, result)
}

func TestMapFuncEmpty(t *testing.T) {
	as := assert.New(t)
	arr := []int{}
	f := func(x int) int { return x * 2 }
	result := MapFunc(arr, f)
	as.Equal([]int{}, result)
}

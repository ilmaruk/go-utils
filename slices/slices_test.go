package slices_test

import (
	"testing"

	"github.com/ilmaruk/go-utils/slices"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	items := []int{1, 5, 2, 4, 3}
	filtered := slices.Filter(items, func(i int) bool { return i > 3 })
	require.ElementsMatch(t, []int{4, 5}, filtered)
}

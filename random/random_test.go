package random_test

import (
	"math/rand"
	"testing"

	"github.com/ilmaruk/go-utils/random"
	"github.com/stretchr/testify/require"
)

func TestChoice(t *testing.T) {
	testCases := []struct {
		name     string
		items    []int
		expected int
	}{
		{
			name:     "empty slice",
			items:    []int{},
			expected: *new(int),
		},
		{
			name:     "single item slice",
			items:    []int{1},
			expected: 1,
		},
		{
			name:     "non empty slice",
			items:    []int{1, 2, 3, 4, 5},
			expected: 5,
		},
	}

	zeroRand := rand.New(rand.NewSource(0))

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			random.SetRand(zeroRand)
			val := random.Choice(tc.items)
			require.Equal(t, tc.expected, val)
		})
	}
}

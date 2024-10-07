package random_test

import (
	"testing"

	"github.com/ilmaruk/go-utils/random"
	"github.com/stretchr/testify/require"
)

func TestRand_Choice(t *testing.T) {
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

	rnd := random.NewWithSeed[int](0)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val := rnd.Choice(tc.items)
			require.Equal(t, tc.expected, val)
		})
	}
}

func TestRand_Shuffle(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rnd := random.NewWithSeed[int](0)
	output := rnd.Shuffle(input)
	require.ElementsMatch(t, input, output)
	require.NotEqual(t, input, output)
}

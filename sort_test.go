package counting

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		description string
		input       []int
		max         int
		expected    []int
	}{
		{
			"Simple",
			[]int{1, 4, 1, 2, 7, 5, 2},
			9,
			[]int{1, 1, 2, 2, 4, 5, 7},
		},
	}

	for _, test := range tests {
		actual := Sort(test.input, test.max)

		for i := 0; i < len(actual); i++ {
			if actual[i] != test.expected[i] {
				t.Log(actual)
				t.Log(test.expected)
				t.Fatal()
			}
		}
	}
}

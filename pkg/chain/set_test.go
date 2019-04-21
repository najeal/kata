package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetIsIn(t *testing.T) {
	tests := []struct {
		name     string
		setInit  map[string]bool
		search   string
		expected bool
	}{
		{
			"should be present",
			map[string]bool{"work": true},
			"work",
			true,
		},
		{
			"should not be present",
			map[string]bool{"work": true},
			"newwork",
			false,
		},
	}

	for _, utest := range tests {
		t.Run(utest.name, func(t *testing.T) {
			set := &Set{set: utest.setInit}
			res := set.IsIn(utest.search)
			assert.Equal(t, utest.expected, res)
		})
	}
}

func TestSetAddAndGetValues(t *testing.T) {
	tests := []struct {
		name        string
		add         []string
		expectedLen int
	}{
		{
			"add one",
			[]string{"work"},
			1,
		},
		{
			"add two",
			[]string{"work", "newwork"},
			2,
		},
		{
			"add three",
			[]string{"work", "newwork", "withoutwork"},
			3,
		},
	}

	for _, utest := range tests {
		t.Run(utest.name, func(t *testing.T) {
			set := NewSet()
			for _, elem := range utest.add {
				set.Add(elem)
				isIn := set.IsIn(elem)
				assert.Equal(t, true, isIn)
			}
			assert.Equal(t, utest.expectedLen, len(set.set))
			values := set.GetValues()
			assert.ElementsMatch(t, utest.add, values)
		})
	}
}

package usecase

import (
	"reflect"
	"testing"
)

func TestGetDifference(t *testing.T) {
	scheduler := New(nil, nil, nil, nil, nil, nil)
	cases := []struct {
		first  []string
		second []string
		differ []string
	}{
		{
			first:  []string{"a", "b", "c"},
			second: []string{"a", "b"},
			differ: []string{"c"},
		},
		{
			first:  []string{},
			second: []string{"a", "b"},
			differ: []string{},
		},
		{
			first:  []string{"a", "b", "c"},
			second: []string{},
			differ: []string{"a", "b", "c"},
		},
		{
			first:  []string{},
			second: []string{},
			differ: []string{},
		},
	}

	for _, c := range cases {
		actual := scheduler.getDifference(c.first, c.second)

		if len(actual) != len(c.differ) && !reflect.DeepEqual(actual, c.differ) {
			t.Errorf("Expected %v, got %v", c.differ, actual)
		}

	}
}

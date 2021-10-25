package cron

import (
	"reflect"
	"strings"
	"testing"
)

func Test_ParseExpression(t *testing.T) {
	p := new(Parser)
	expr := "*/15 0 1,15 * 1-5 /usr/bin/find"
	wants := Entry{
		Minutes:    field{0, 15, 30, 45},
		Hour:       field{0},
		DayOfMonth: field{1, 15},
		Month:      field{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		DayOfWeek:  field{1, 2, 3, 4, 5},
		Command:    "/usr/bin/find",
	}
	got, err := p.Parse(expr)
	if err != nil {
		t.Errorf("%s => unexpected error %v", expr, err)
	}
	if !reflect.DeepEqual(wants, got) {
		t.Errorf("%s => wants %v, got %v", expr, wants, got)
	}
}

func Test__ParseRange(t *testing.T) {
	empty := make([]uint, 0)
	tc := []struct {
		expr     string
		min, max uint
		wants    []uint
		err      string
	}{
		{"*/15", 0, 59, []uint{0, 15, 30, 45}, ""},
		{"0", 0, 23, []uint{0}, ""},
		{"1-5", 1, 31, []uint{1, 2, 3, 4, 5}, ""},
		{"*", 0, 6, []uint{0, 1, 2, 3, 4, 5, 6}, ""},
		{"1,5", 0, 0, empty, "failed to parse int"},
		{"1-5/2/3", 0, 0, empty, "invalid step range"},
		{"*/-12", 0, 6, empty, "negative number"},
		{"1", 3, 6, empty, "below minimum"},
		{"6", 3, 5, empty, "above maximum"},
		{"5-3", 3, 5, empty, "beyond end of range"},
		{"*/0", 0, 0, empty, "should be a positive number"},
	}

	for _, c := range tc {
		got, err := ParseRange(c.expr, bounds{min: c.min, max: c.max})
		if len(c.err) != 0 && (err == nil || !strings.Contains(err.Error(), c.err)) {
			t.Errorf("%s => expected %v, got %v", c.expr, c.err, err)
		}
		if len(c.err) == 0 && err != nil {
			t.Errorf("%s => unexpected error %v", c.expr, err)
		}
		if !reflect.DeepEqual(c.wants, got) {
			t.Errorf("%s => wants %v, got %v", c.expr, c.wants, got)
		}
	}
}

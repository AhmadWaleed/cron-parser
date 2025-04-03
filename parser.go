package cron

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MaxAllowedFields = 6
	DefaultExpr      = "0 0 0 * * *"
)

type Parser struct{ expr string }

// NewDefaultParser creates new parser with default cron expression.
func NewDefaultParser(cmd string) *Parser {
	p := new(Parser)
	p.expr = fmt.Sprintf("%s %s", DefaultExpr, cmd)
	return p
}

// Parse parses a cron string and expands each field to and
// return the parsed entry which can be used for any purpose
// or simply for printing parsed cron expression in tabular format.
func (p Parser) Parse(spec string) (Entry, error) {
	entry := Entry{}

	if len(spec) == 0 {
		return entry, fmt.Errorf("cron spec string cannot be empty.")
	}

	// Split on whitespaces.
	fields := strings.Fields(spec)
	if len(fields) < MaxAllowedFields {
		return entry, fmt.Errorf("expected exactly %d fields, found %d, %s", MaxAllowedFields, len(fields), spec)
	}

	entry.Minutes = MustParseField(fields[0], minutes)
	entry.Hour = MustParseField(fields[1], hours)
	entry.DayOfMonth = MustParseField(fields[2], dom)
	entry.Month = MustParseField(fields[3], months)
	entry.DayOfWeek = MustParseField(fields[4], dow)
	entry.Command = fields[5]

	return entry, nil
}

// MustParseField parse the field expression and
// panic out in case of error.
func MustParseField(field string, b bounds) []uint {
	v, err := ParseField(field, b)
	if err != nil {
		panic(err)
	}
	return v
}

// ParseField split expression field into command-separated form
// and then parse the range for each field.
func ParseField(field string, b bounds) ([]uint, error) {
	var result []uint
	ranges := strings.FieldsFunc(field, func(r rune) bool {
		return r == ','
	})

	var err error
	for _, expr := range ranges {
		val, err := ParseRange(expr, b)
		if err == nil {
			result = append(result, val...)
		}
	}

	return result, err
}

// ParseRange parse the field expression by range and step:
// number | number "-" number [ "/" number ]
// or error parsing range.
func ParseRange(expr string, b bounds) ([]uint, error) {
	var (
		min, max, step uint
		rangeAndStep   = strings.Split(expr, "/")
		lowAndHigh     = strings.Split(rangeAndStep[0], "-")
		err            error
	)

	if lowAndHigh[0] == "*" {
		min = b.min
		max = b.max
	} else {
		min, err = ParseStrOrName(lowAndHigh[0], b.dict)

		switch len(lowAndHigh) {
		case 1:
			max = min
		case 2:
			max, err = ParseStrOrName(lowAndHigh[1], b.dict)
		default:
			err = fmt.Errorf("invalid range given: %s", expr)
		}
	}

	result := make([]uint, 0)
	if err != nil {
		return result, err
	}

	switch len(rangeAndStep) {
	case 1:
		step = 1
	case 2:
		step, err = ParseStrOrName(rangeAndStep[1], b.dict)
	default:
		err = fmt.Errorf("invalid step range: %s", expr)
	}

	if err != nil {
		return result, err
	}

	err = isValidRange(min, max, step, b)
	if err != nil {
		return result, err
	}

	for i := min; i <= max; i += step {
		result = append(result, i)
	}

	return result, err
}

func ParseStrOrName(expr string, dict map[string]uint) (uint, error) {
	if dict != nil {
		if val, ok := dict[strings.ToLower(expr)]; ok {
			return val, nil
		}
	}

	return MustParseInt(expr)
}

func MustParseInt(expr string) (uint, error) {
	num, err := strconv.Atoi(expr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse int from %s: %s", expr, err)
	}
	if num < 0 {
		return 0, fmt.Errorf("negative number (%d) not allowed: %s", num, expr)
	}

	return uint(num), nil
}

// isValidRange validates the range bounds.
func isValidRange(min, max, step uint, b bounds) error {
	if min < b.min {
		return fmt.Errorf("beginning of range (%d) below minimum (%d)", min, b.min)
	}
	if max > b.max {
		return fmt.Errorf("end of range (%d) above maximum (%d)", max, b.max)
	}
	if min > max {
		return fmt.Errorf("beginning of range (%d) beyond end of range (%d)", min, max)
	}
	if step == 0 {
		return fmt.Errorf("step of range should be a positive number")
	}
	return nil
}

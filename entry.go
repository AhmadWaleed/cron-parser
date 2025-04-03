package cron

import (
	"fmt"
	"strconv"
	"strings"
)

type field []uint

func (nums field) string() string {
	var s []string
	for _, n := range nums {
		s = append(s, strconv.Itoa(int(n)))
	}
	return strings.Join(s, " ")
}

// Entry represents the parsed cron expresssion.
type Entry struct {
	Minutes    field
	Hour       field
	DayOfMonth field
	Month      field
	DayOfWeek  field
	Command    string
}

// Print the parsed CronExpr in tabular format.
func (e Entry) Print() {
	fmt.Printf(
		"minute\t\t%s\nhour\t\t%s\nday of month\t%s\nmonth\t\t%s\nday of week\t%s\ncommand\t\t%s\n",
		e.Minutes.string(), e.Hour.string(), e.DayOfMonth.string(), e.Month.string(), e.DayOfWeek.string(), e.Command)
}

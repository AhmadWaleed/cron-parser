// Package cron provides a parser for standard cron expressions and a utility to print
// the parsed fields in a tabular format. The package follows the common cron syntax used
// in Unix-based systems, allowing users to analyze and understand cron schedules easily.
//
// Cron Expression Format:
// The parser supports the five-field cron format followed by an optional command:
//
//   - *    *    *    *    command
//     ┬    ┬    ┬    ┬    ┬
//     │    │    │    │    │
//     │    │    │    │    └─ Day of the week (0-6 or SUN-SAT)
//     │    │    │    └────── Month (1-12 or JAN-DEC)
//     │    │    └────────── Day of the month (1-31)
//     │    └─────────────── Hour (0-23)
//     └──────────────────── Minute (0-59)
//
// Supported Field Types:
// - Wildcards (*) match all possible values.
// - Single values (e.g., "5") specify a fixed time unit.
// - Lists (e.g., "1,2,3") allow multiple specific values.
// - Ranges (e.g., "1-5") define a range of values.
// - Steps (e.g., "*/5") specify intervals within a range.
//
// Example Usage:
//
//	p := new(cron.Parser)
//	expr := "*/15 0 1,15 * 1-5 /usr/bin/find"
//	entry, err := p.Parse(expr)
//	if err != nil {
//	    fmt.Fprintf(os.Stderr, "%s", err)
//	}
//
//	entry.Print()
//
// The output will be a structured representation of the parsed fields along with the command.
//
// This package is useful for debugging cron expressions and understanding scheduled tasks.
package cron

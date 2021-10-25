package main

import (
	"fmt"
	"os"

	"github.com/ahmadwaleed/cron-parser"
)

func main() {
	p := new(cron.Parser)
	expr := "*/15 0 1,15 * 1-5 /usr/bin/find"
	entry, err := p.Parse(expr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}

	entry.Print()
}
// Example Output

// minute          0 15 30 45
// hour            0
// day of month    1 15
// month           1 2 3 4 5 6 7 8 9 10 11 12
// day of week     1 2 3 4 5
// command         /usr/bin/find

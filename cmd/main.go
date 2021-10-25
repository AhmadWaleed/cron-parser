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

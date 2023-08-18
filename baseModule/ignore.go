package main

import (
	"os"

	"golang.org/x/example/hello/reverse"
	"golang.org/x/net/html"
)

func thisCallsVulnerableCode(s string) string {
	html.Parse(os.Stdin)
	return reverse.String(s)
}

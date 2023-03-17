package main

import (
	"termcheck"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(termcheck.Analyzer) }

package main

import (
	"github.com/Keyl0ve/termcheck"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(termcheck.Analyzer) }

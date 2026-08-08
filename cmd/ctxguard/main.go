package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"lesiw.io/ctxguard"
)

func main() { singlechecker.Main(ctxguard.Analyzer) }

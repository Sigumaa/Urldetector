package main

import (
	"github.com/Sigumaa/Urldetector"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(Urldetector.Analyzer) }

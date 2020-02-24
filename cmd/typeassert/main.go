package main

import (
	"github.com/gostaticanalysis/typeassert"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(typeassert.Analyzer) }
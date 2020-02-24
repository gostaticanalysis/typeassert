package typeassert_test

import (
	"testing"

	"github.com/gostaticanalysis/typeassert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, typeassert.Analyzer, "a")
}
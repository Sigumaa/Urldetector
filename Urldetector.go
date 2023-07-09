package Urldetector

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "Urldetector is a tool to detect identifiers containing 'Url'."

// Analyzer is an analysis analyzer that finds identifiers containing 'Url'.
var Analyzer = &analysis.Analyzer{
	Name: "Urldetector",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			if n.Name == "Url" {
				pass.Reportf(n.Pos(), "Url should be URL")
			}
		}
	})

	return nil, nil
}

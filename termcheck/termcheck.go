package termcheck

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "termcheck is the linter to check using simple term"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "termcheck",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.SelectorExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.SelectorExpr:
			// user.Read などのフィールドやメソッドを参照する式
			getSelectorName(pass, n)
		}
	})

	return nil, nil
}

func getSelectorName(pass *analysis.Pass, selectorExpr *ast.SelectorExpr) {
	// selector の左の名前を取得する ->  X
	leftExpr := selectorExpr.X
	leftIdent, ok := leftExpr.(*ast.Ident)
	if !ok || len(leftIdent.Name) == 0 {
		return
	}

	// selector の右の名前を取得する ->  Sel
	rightIdentName := selectorExpr.Sel.Name
	if len(rightIdentName) == 0 {
		return
	}

	// 同じ term が使われているかどうかを判断する
	if ok := isContainsDuplicate(leftIdent.Name, rightIdentName); !ok {
		return
	}
	pass.Reportf(selectorExpr.Pos(), "word is used multiple in same line")
}

func isContainsDuplicate(leftName, rightName string) bool {
	// 左の文字が 1,2 の時はスキップする
	// TODO:
	if len(leftName) == 1 || len(leftName) == 2 {
		return false
	}
	leftStr := strings.ToLower(leftName)
	rightStr := strings.ToLower(rightName)

	for _, l := range leftStr {
		if strings.ContainsRune(rightStr, l) {
			return true
		}
	}
	return false
}

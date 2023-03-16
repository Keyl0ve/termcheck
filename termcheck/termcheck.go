package termcheck

import (
	"go/ast"
	"strings"

	"github.com/iancoleman/strcase"
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
			leftName, rightName, ok := getSelectorName(pass, n)
			if !ok {
				return
			}

			if !isContainsDuplicate(leftName, rightName) {
				return
			}
			pass.Reportf(n.Pos(), "%s is used multiple in same line", leftName)
		}
	})

	return nil, nil
}

func getSelectorName(pass *analysis.Pass, selectorExpr *ast.SelectorExpr) (string, string, bool) {
	// selector の左の名前を取得する ->  X
	leftExpr := selectorExpr.X
	leftIdent, ok := leftExpr.(*ast.Ident)
	if !ok || len(leftIdent.Name) == 0 {
		return "", "", false
	}

	// selector の右の名前を取得する ->  Sel
	rightIdentName := selectorExpr.Sel.Name
	if len(rightIdentName) == 0 {
		return "", "", false
	}

	return leftIdent.Name, rightIdentName, true
}

func isContainsDuplicate(leftName, rightName string) bool {
	// 左の文字数が 1,2,3 の時はスキップする
	if len(leftName) < 3 {
		return false
	}

	// ReadUserFromJapan -> read_user_from_japan
	targetStr := strcase.ToSnake(rightName)
	return strings.Contains(targetStr, leftName)
}

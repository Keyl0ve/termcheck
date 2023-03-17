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
			rowSlice := appendName(n)
			nameSlice := splitName(rowSlice)

			str, ok := isContainsDuplicate(nameSlice)
			if !ok {
				return
			}
			pass.Reportf(n.Pos(), "%s is used multiple in same line", str)
		}
	})

	return nil, nil
}

func appendName(selectorExpr *ast.SelectorExpr) []string {
	res := []string{}

	if leftIdent, ok := selectorExpr.X.(*ast.Ident); ok {
		res = append(res, leftIdent.Name)
	} else if selector, ok := selectorExpr.X.(*ast.SelectorExpr); ok {
		res = appendName(selector)
	} else {
		panic("unexpected selector")
	}

	res = append(res, selectorExpr.Sel.Name)

	return res
}

func splitName(nameSlice []string) []string {
	strSlice := []string{}
	for _, name := range nameSlice {
		// ReadUserFromJapan -> read_user_from_japan
		snakedStr := strcase.ToSnake(name)
		// read_user_from_japan -> [read, user, from, japan]
		targetStr := strings.Split(snakedStr, "_")

		for _, v := range targetStr {
			strSlice = append(strSlice, v)
		}
	}

	return strSlice
}

func isContainsDuplicate(strSlice []string) (string, bool) {
	// 要素数が1以下の場合は、重複する要素がないので即座にfalseを返します。
	if len(strSlice) < 2 {
		return "", false
	}

	encountered := map[string]bool{}

	for _, str := range strSlice {
		// u, uu など3文字未満の文字の場合 continue
		if len(str) < 3 {
			continue
		}
		if encountered[str] {
			return str, true
		} else {
			encountered[str] = true
		}
	}

	return "", false
}

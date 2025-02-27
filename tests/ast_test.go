package tests

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"testing"
)

func TestAST(t *testing.T) {
	filePath := "../src/router/module.go"
	set := token.NewFileSet()
	content, _ := os.ReadFile(filePath)

	node, err := parser.ParseFile(set, filePath, content, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}

	// TODO 动态解析代码，找到 InitModuleRouter 函数，然后在函数体中添加新的代码
	ast.Inspect(node, func(n ast.Node) bool {
		// 查找函数声明节点
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Name.Name == "InitModuleRouter" {
			// 创建新的语句
			newStmt := &ast.ExprStmt{
				X: &ast.CallExpr{
					Fun:  ast.NewIdent("fmt.Println"),
					Args: []ast.Expr{ast.NewIdent("\"New code added to InitModuleRouter\"")},
				},
			}
			// 将新语句追加到函数体中
			fn.Body.List = append(fn.Body.List, newStmt)

			// 打印修改后的代码
			var buf bytes.Buffer
			printer.Fprint(&buf, set, node)
			os.WriteFile(filePath, buf.Bytes(), 0644) // webssh.InitRouter(app, backendRouter)
		}

		return true
	})
}

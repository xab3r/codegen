package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"reflect"
	"strings"
)

// go build gen/* && ./codegen

func main() {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, "main.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	var out = new(bytes.Buffer)

	fmt.Fprintf(out, "package %s\n", node.Name.Name)
	fmt.Fprintln(out,
		`import (
			"bytes"
			"fmt"
		)`,
	)
	fmt.Fprintln(out)

	for _, decl := range node.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok {
			fmt.Printf("Skipping %v - it's not *ast.GenDecl\n", decl)
			continue
		}

	SPECS:
		for _, spec := range gen.Specs {
			givenType, ok := spec.(*ast.TypeSpec)
			if !ok {
				fmt.Printf("Skipping %v - it's not *ast.TypeSpec\n", spec)
				continue
			}

			givenStruct, ok := givenType.Type.(*ast.StructType)
			if !ok {
				fmt.Printf("Skipping %v - it's not *ast.StructType\n", givenType)
				continue SPECS
			}

			fmt.Fprintf(out,
				`func PureFormat%s(s %s) string {
					var buf = &bytes.Buffer{}

				`,
				givenType.Name.Name,
				givenType.Name.Name,
			)

		FILEDS:
			for _, field := range givenStruct.Fields.List {
				if field.Tag != nil {
					tag := reflect.StructTag(field.Tag.Value)

					if tag.Get("print") != "true" {
						continue FILEDS
					}

					fieldName := field.Names[0].Name
					fmt.Fprintf(
						out,
						`fmt.Fprintf(buf, "%s: %%v\n", s.%s)`,
						strings.ToLower(fieldName),
						fieldName,
					)
					fmt.Fprintln(out)

				}
			}

			fmt.Fprintln(out,
				`
					return buf.String()
				}`,
			)

		}
	}

	srcSet := token.NewFileSet()
	srcAstFile, err := parser.ParseFile(srcSet, "", out, 0)
	if err != nil {
		panic("Something went wrong")
	}

	src, _ := os.Create("pure.go")
	defer src.Close()

	fmt.Fprintf(src, "// DO NOT EDIT\n\n")
	printer.Fprint(src, srcSet, srcAstFile)
}

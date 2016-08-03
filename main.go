package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func handle(list []*ast.Comment) {
	active := false
	for _, c := range list {
		text := c.Text
		if strings.HasPrefix(text, "//") {
			text = c.Text[2:]
		}
		text = strings.Trim(text, " ") // some comments do have a space after the slashes, some don't
		parts := strings.Split(text, " ")
		if len(parts) > 3 && parts[0] == "metric" && parts[2] == "is" {
			active = true
			fmt.Printf("* `%s`:  \n", parts[1])
			fmt.Println(text[len(parts[0])+len(parts[1])+len(parts[2])+3:])
		} else if active {
			fmt.Println(text)
		}
	}
}

func main() {
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		log.Fatal("metrics2docs <path-to-codebase>")
	}
	searchDir := os.Args[1]
	mode := parser.ParseComments
	fset := token.NewFileSet() // positions are relative to fset
	fmt.Println("# overview of metrics")
	fmt.Printf("(only shows metrics that are documented. generated with [metrics2docs](github.com/Dieterbe/metrics2docs))\n\n")
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		if !strings.HasSuffix(f.Name(), ".go") {
			return nil
		}
		src, err := parser.ParseFile(fset, path, nil, mode)
		if err != nil {
			return err
		}
		for _, d := range src.Decls {
			if d, ok := d.(*ast.GenDecl); ok {
				for _, s := range d.Specs {
					if s, ok := s.(*ast.ValueSpec); ok {
						if s.Doc != nil {
							handle(s.Doc.List)
						}
						if s.Comment != nil {
							handle(s.Comment.List)
						}
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

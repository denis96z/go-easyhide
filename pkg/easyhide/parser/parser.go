package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type Parser struct {
	fname string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseFile(fname string) (Info, error) {
	fset := token.NewFileSet()

	astf, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		return Info{}, fmt.Errorf("failed to parse file [fname = %q]: %w", fname, err)
	}

	//ast.Print(fset, astf)

	v := newVisitor()
	ast.Walk(v, astf)

	v.info.Pkg.Path, err = getPkgPath(fname)
	if err != nil {
		return Info{}, fmt.Errorf("failed to get package path [fname = %q]: %w", fname, err)
	}

	return v.info, nil
}

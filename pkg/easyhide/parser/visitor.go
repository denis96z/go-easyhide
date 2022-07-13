package parser

import (
	"go/ast"
)

type visitor struct {
	info Info
}

func newVisitor() *visitor {
	return &visitor{
		info: Info{
			Types: make([]TypeInfo, 0),
		},
	}
}

func (v *visitor) Visit(n ast.Node) (w ast.Visitor) {
	switch n := n.(type) {
	case *ast.Package:
		return v

	case *ast.File:
		v.info.Pkg.Name = n.Name.Name
		return v

	case *ast.GenDecl:
		if n.Doc == nil {
			return nil
		}
		for _, cmt := range n.Doc.List {
			txt := GetCommentText(cmt.Text)
			if txt != "easyhide:json" {
				return nil
			}
		}
		return v

	case *ast.TypeSpec:
		v.info.Types = append(
			v.info.Types,
			TypeInfo{
				Name: n.Name.Name,
			},
		)
		return v

	case *ast.StructType:
		return v
	}
	return nil
}

package generator

const (
	FileTemplate = `
{{$data := .}}
package {{$data.PkgName}}

// Code generated by easyhide. DO NOT EDIT.

//go:generate easyjson

import (
	json "encoding/json"

	easyhide "github.com/denis96z/go-easyhide/pkg/easyhide"
)

{{range $idx, $tp := $data.Types}}
	{{$tp.TypeCode}}

	{{$tp.FuncCode}}
{{end}}
`
)

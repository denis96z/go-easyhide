package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/denis96z/go-easyhide/pkg/easyhide/generator"
	"github.com/denis96z/go-easyhide/pkg/easyhide/parser"
)

func main() {
	fname := os.Getenv("GOFILE")
	if !strings.HasSuffix(fname, ".go") {
		log.Fatalf("not a go file [name = %q]", fname)
	}

	afname, err := filepath.Abs(fname)
	if err != nil {
		log.Fatalf("failed to get absolute file name [name = %q]: %s", fname, err)
	}

	p := parser.NewParser()

	info, err := p.ParseFile(afname)
	if err != nil {
		log.Fatalf("failed to parse file [name = %q]: %s", afname, err)
	}

	info.GeneratedFilePath = generator.MakeGeneratedFilePath(afname)

	gfpath := generator.MakeGeneratorFilePath(afname)
	if err = generator.WriteGeneratorFile(gfpath, info); err != nil {
		log.Fatalf("failed to write generator file [path = %q]: %s", gfpath, err)
	}
	if err = generator.GoRunGeneratorFile(gfpath); err != nil {
		log.Fatalf("failed to run generator file [path = %q]: %s", gfpath, err)
	}
	if err = generator.RemoveGeneratorFile(gfpath); err != nil {
		log.Fatalf("failed to remove generator file [path = %q]: %s", gfpath, err)
	}
}

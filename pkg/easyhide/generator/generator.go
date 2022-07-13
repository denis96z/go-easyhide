package generator

import (
	"fmt"
	"os"

	"github.com/denis96z/go-easyhide/pkg/easyhide/parser"
)

func MakeGeneratorFilePath(srcPath string) string {
	bname := srcPath[:len(srcPath)-len(".go")]
	return bname + "_easyhide_generator.go"
}

func WriteGeneratorFile(fpath string, info parser.Info) error {
	if err := createGoFileFromTemplate(fpath, "generator", GenFileTemplate, info); err != nil {
		return fmt.Errorf("failed to write generator file [path = %q]: %w", fpath, err)
	}
	return nil
}

func RemoveGeneratorFile(fpath string) error {
	if err := os.Remove(fpath); err != nil {
		return fmt.Errorf("failed to remove generator file [path = %q]: %w", fpath, err)
	}
	return nil
}

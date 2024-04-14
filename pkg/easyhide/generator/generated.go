package generator

import (
	"fmt"
)

func MakeGeneratedFilePath(srcPath string) string {
	bname := srcPath[:len(srcPath)-len(".go")]
	return bname + "_easyhide.go"
}

func GoRunGeneratorFile(gfpath string) error {
	return executeCommand("go", "run", gfpath)
}

func GoRunGenerateForResultFile(path string) error {
	return executeCommand("go", "generate", path)
}

type GenData struct {
	PkgName string

	EasyJSON bool

	Types []GenDataTypeInfo
}

type GenDataTypeInfo struct {
	TypeCode string
	FuncCode string
}

func WriteEasyHideFile(fpath string, gData GenData) error {
	if err := createGoFileFromTemplate(fpath, "easyhide", FileTemplate, gData); err != nil {
		return fmt.Errorf("failed to write easyhide file [path = %q]: %w", fpath, err)
	}
	return nil
}

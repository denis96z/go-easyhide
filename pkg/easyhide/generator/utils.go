package generator

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"text/template"
)

func GenerateCodeForType(v any, tpName string) (string, string) {
	xtp := "tp" + tpSHA256(tpName)

	s1 := "//easyjson:json\ntype " + xtp + " " + tpName + "\n"
	s2 := "func (v " + tpName + ") EasyHide() ([]byte, error) {\n"

	k := reflect.ValueOf(v).Kind()
	switch k {
	case reflect.Struct:
		ss2 := generateCodeForStruct(v, xtp)
		s2 += ss2

	default:
		panic("bad type")
	}

	s2 += "}"
	return s1, s2
}

func generateCodeForStruct(v any, tp string) string {
	s := "xv := " + tp + "(v)\n"
	s += "if easyhide.HideData {\n"

	val := reflect.ValueOf(v)
	s += generateCodeForStructFields(val, "xv.")

	s += "}\nreturn json.Marshal(xv)\n"
	return s
}

func generateCodeForStructFields(val reflect.Value, pfx string) string {
	s := ""
	for i := 0; i < val.NumField(); i++ {
		f := val.Type().Field(i)
		switch f.Type.Kind() {
		case reflect.Struct:
			s += generateCodeForStructFields(val.Field(i), pfx+f.Name+".")

		case reflect.String:
			tg := f.Tag.Get("easyhide")
			if tg == "" || tg == "show" {
				continue
			}

			s += generateCodeForString(pfx+f.Name, tg) + "\n"
		}
	}
	return s
}

func generateCodeForString(name string, tg string) string {
	switch {
	case tg == "hide":
		return name + " = easyhide.HiddenMarker"
	case tg == "hide:HL":
		return name + " = easyhide.HiddenMarker + " + name + "[len(" + name + ")/2:]"
	case tg == "hide:HR":
		return name + " = " + name + "[:len(" + name + ")/2] + easyhide.HiddenMarker"
	case tg == "hide:NE":
		return `if ` + name + ` != "" { ` + name + ` = easyhide.HiddenMarker }`
	case tg == "hide:HL,NE":
		return `if ` + name + ` != "" { ` + name + " = easyhide.HiddenMarker + " + name + "[len(" + name + ")/2:] }"
	case tg == "hide:HR,NE":
		return `if ` + name + ` != "" { ` + name + " = " + name + "[:len(" + name + ")/2] + easyhide.HiddenMarker }"
	case strings.HasPrefix(tg, "hide:RE"):
		rst, ne := tg[len("hide:RE"):], false
		if strings.HasPrefix(rst, ",NE") {
			rst, ne = rst[len(",NE"):], true
		}
		if !strings.HasPrefix(rst, ":") {
			panic(fmt.Sprintf(
				"bad tag rest: %q", rst),
			)
		}
		rst = rst[len(":"):]
		spl := strings.Split(rst, ":")
		if len(spl) != 2 {
			panic(fmt.Sprintf(
				"bad tag rest: %q", rst),
			)
		}
		rxpName, rplName := spl[0], spl[1]
		if rxpName == "" {
			panic(errors.New("empty regexp name"))
		}
		if rplName == "" {
			panic(errors.New("empty replacement name"))
		}
		s := ""
		if ne {
			s += `if ` + name + ` != "" { `
		}
		s += name + ` = ` + rxpName + `.ReplaceAllString(` + name + `, ` + rplName + `)`
		if ne {
			s += ` }`
		}
		return s
	default:
		panic(fmt.Sprintf(
			"bad tag: %q", tg,
		))
	}
}

func tpSHA256(tp string) string {
	h := sha256.New()
	h.Write([]byte(tp))
	return hex.EncodeToString(
		h.Sum(nil),
	)
}

func createGoFileFromTemplate(fpath string, name, ttxt string, data any) error {
	tmpl, err := template.New(name).Parse(ttxt)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err = (tmpl).Execute(&buf, data); err != nil {
		return fmt.Errorf("template %q execute error: %w", name, err)
	}

	fset := token.NewFileSet()

	astf, err := parser.ParseFile(fset, fpath, buf.Bytes(), parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse file content: %w", err)
	}

	f, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf(
			"failed to create file: %w", err,
		)
	}
	defer func() {
		_ = f.Close()
	}()
	if err = format.Node(f, fset, astf); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func executeCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	{
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			`failed to run command: %s`, err,
		)
	}
	return nil
}

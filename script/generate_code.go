package script

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateCode(dirName string) {
	funcName := toFuncName(dirName)
	packageName := strings.ReplaceAll(dirName, "-", "_")
	dirPath := fmt.Sprintf("./%s", dirName)
	os.MkdirAll(dirPath, os.ModePerm)

	writeFile := func(filename, tmpl string) {
		path := fmt.Sprintf("%s/%s", dirPath, filename)
		f, _ := os.Create(path)
		defer f.Close()
		t := template.Must(template.New("tmpl").Parse(tmpl))
		t.Execute(f, map[string]string{
			"PackageName": packageName,
			"FuncName":    funcName,
		})
		fmt.Println("✅ Created:", path)
	}

	writeFile(fmt.Sprintf("%s.go", packageName), codeTemplate)
	writeFile(fmt.Sprintf("%s_test.go", packageName), testTemplate)
}

func toFuncName(slug string) string {
	caser := cases.Title(language.English)
	parts := strings.Split(slug, "-")
	for i, p := range parts {
		parts[i] = caser.String(p)
	}
	return strings.Join(parts, "")
}

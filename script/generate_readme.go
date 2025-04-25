package script

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateReadme() {
	files, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}

	var lines []string
	lines = append(lines, "# LeetCode Solutions\n")

	caser := cases.Title(language.English)

	for _, f := range files {
		if f.IsDir() && !strings.HasPrefix(f.Name(), ".") && f.Name() != "script" {

			title := strings.ReplaceAll(f.Name(), "-", " ")
			title = caser.String(title)

			link := fmt.Sprintf("- [%s](%s)", title, f.Name())
			lines = append(lines, link)
		}
	}

	err = os.WriteFile("./README.md", []byte(strings.Join(lines, "\n")+"\n"), 0o644)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ README.md updated.")
}

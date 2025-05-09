package script

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	codeFileSuffix          = ".go"
	testFileSuffix          = "_test.go"
	readmeFile              = "README.md"
	leetcodeSolutionsHeader = "# LeetCode Solutions"
)

// Generate creates a new LeetCode solution with slug
func Generate(slug string) {
	if slug == "" {
		panic("slug cannot be empty")
	}

	var nonEmptySlugPartList []string
	for part := range strings.SplitSeq(slug, "-") {
		if part != "" {
			nonEmptySlugPartList = append(nonEmptySlugPartList, strings.ToLower(part))
		} else {
			panic("invalid slug format: empty part found")
		}
	}

	if len(nonEmptySlugPartList) == 0 {
		panic("invalid slug format: no valid parts found")
	}

	dir := strings.Join(nonEmptySlugPartList, "-")
	fileName := strings.Join(nonEmptySlugPartList, "_")
	packageName := fileName

	codeFile := filepath.Join(dir, fileName+codeFileSuffix)
	testFile := filepath.Join(dir, fileName+testFileSuffix)

	pathsToCheck := []string{dir, codeFile, testFile}
	for _, path := range pathsToCheck {
		if _, err := os.Stat(path); err == nil {
			panic(fmt.Sprintf("%s already exists", path))
		} else if !os.IsNotExist(err) {
			panic(fmt.Sprintf("error checking %s: %v", path, err))
		}
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("failed to create directory %s: %v", dir, err))
	}

	if err := writeFile(codeFile, codeTemplate, packageName, nonEmptySlugPartList); err != nil {
		panic(fmt.Sprintf("failed to write code file: %v", err))
	}

	if err := writeFile(testFile, testTemplate, packageName, nonEmptySlugPartList); err != nil {
		panic(fmt.Sprintf("failed to write test file: %v", err))
	}

	if err := updateReadme(slug, codeFile); err != nil {
		panic(fmt.Sprintf("failed to update README: %v", err))
	}
}

// writeFile creates a file with the given template and data
func writeFile(file, tmpl, packageName string, slugPartList []string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", file, err)
	}
	defer f.Close()

	t, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	caser := cases.Title(language.English)

	funcName := ""
	testFuncName := "Test"

	for _, v := range slugPartList {
		v := strings.ToLower(v)
		titleV := caser.String(v)

		if funcName == "" {
			funcName += v
		} else {
			funcName += titleV
		}

		testFuncName += titleV
	}

	data := map[string]string{
		"PackageName":  packageName,
		"FuncName":     funcName,
		"TestFuncName": testFuncName,
	}

	if err := t.Execute(f, data); err != nil {
		return fmt.Errorf("failed to write to file %s: %w", file, err)
	}

	fmt.Println("✅ Created:", file)
	return nil
}

// updateReadme adds the new solution to README.md
func updateReadme(slug, codeFile string) error {
	readmePath := "README.md"
	newEntry := fmt.Sprintf("- [%s](%s)", slug, codeFile)

	contentBytes, readErr := os.ReadFile(readmePath)
	var existingEntries []string

	if readErr == nil {
		lines := strings.SplitSeq(string(contentBytes), "\n")
		for line := range lines {
			if strings.HasPrefix(line, "- [") {
				existingEntries = append(existingEntries, line)
			}
		}
	} else if !os.IsNotExist(readErr) {
		return fmt.Errorf("error reading %s: %w", readmePath, readErr)
	}

	found := slices.Contains(existingEntries, newEntry)

	if !found {
		// Add new entry at the beginning
		updatedEntries := append([]string{newEntry}, existingEntries...)

		var contentLines []string
		contentLines = append(contentLines, leetcodeSolutionsHeader)
		contentLines = append(contentLines, "") // Add an empty line after the header
		contentLines = append(contentLines, updatedEntries...)
		contentLines = append(contentLines, "") // Add an empty line at the end

		updatedContent := strings.Join(contentLines, "\n")

		if err := os.WriteFile(readmePath, []byte(updatedContent), 0o644); err != nil {
			return fmt.Errorf("error updating %s: %w", readmePath, err)
		}

		fmt.Println("✅", readmePath, "updated.")
	} else {
		fmt.Println("✅", readmePath, "already contains entry for", slug)
	}

	return nil
}

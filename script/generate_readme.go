package script

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FileInfo struct {
	Path       string
	Name       string
	IsDir      bool
	CommitTime time.Time
}

func GenerateReadme() {
	lines := []string{"# LeetCode Solutions", ""}

	if err := processDir(".", &lines); err != nil {
		panic(err)
	}

	if err := os.WriteFile("README.md", []byte(strings.Join(lines, "\n")+"\n"), 0o644); err != nil {
		panic(err)
	}

	fmt.Println("✅ README.md updated.")
}

func processDir(dir string, lines *[]string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var fileInfos []FileInfo
	for _, entry := range entries {
		if shouldSkipFile(entry) {
			continue
		}

		entryPath := filepath.Join(dir, entry.Name())

		if entry.IsDir() {
			fileInfos = append(fileInfos, FileInfo{
				Path:  entryPath,
				Name:  entry.Name(),
				IsDir: true,
			})
		} else if strings.HasSuffix(entry.Name(), ".go") && entry.Name() != "main.go" {
			commitTime, err := getGitCommitTime(entryPath)
			if err != nil {
				commitTime = time.Time{}
			}

			fileInfos = append(fileInfos, FileInfo{
				Path:       entryPath,
				Name:       entry.Name(),
				IsDir:      false,
				CommitTime: commitTime,
			})
		}
	}

	sortFileInfosByGit(fileInfos)

	caser := cases.Title(language.English)
	for _, info := range fileInfos {
		if info.IsDir {
			if err := processDir(info.Path, lines); err != nil {
				return err
			}
		} else {
			relativePath := buildRelativePath(dir, info.Name)
			title := formatTitle(info.Name, caser)
			link := fmt.Sprintf("- [%s](%s)", title, relativePath)
			*lines = append(*lines, link)
		}
	}

	return nil
}

func getGitCommitTime(filePath string) (time.Time, error) {
	cmd := exec.Command("git", "log", "-1", "--format=%ct", filePath)
	output, err := cmd.Output()
	if err != nil {
		return time.Time{}, err
	}

	timestamp := strings.TrimSpace(string(output))
	if timestamp == "" {
		return time.Time{}, fmt.Errorf("no git history for file: %s", filePath)
	}

	var unixTime int64
	if _, err := fmt.Sscanf(timestamp, "%d", &unixTime); err != nil {
		return time.Time{}, err
	}

	return time.Unix(unixTime, 0), nil
}

func sortFileInfosByGit(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		if !files[i].CommitTime.Equal(files[j].CommitTime) {
			return files[i].CommitTime.After(files[j].CommitTime)
		}
		return files[i].Name < files[j].Name
	})
}

func shouldSkipFile(f os.DirEntry) bool {
	if f.IsDir() {
		return strings.HasPrefix(f.Name(), ".") || f.Name() == "script"
	}
	if strings.HasSuffix(f.Name(), ".go") {
		return f.Name() == "main.go" || strings.HasSuffix(f.Name(), "_test.go")
	}
	return true
}

func buildRelativePath(dir, fileName string) string {
	relativePath := filepath.Join(dir, fileName)
	relativePath = strings.TrimPrefix(relativePath, "./")
	return filepath.ToSlash(relativePath)
}

func formatTitle(fileName string, caser cases.Caser) string {
	title := strings.ReplaceAll(fileName, "_", " ")
	title = strings.TrimSuffix(title, ".go")
	return caser.String(title)
}

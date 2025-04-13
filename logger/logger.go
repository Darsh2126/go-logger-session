package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func WriteSession(start, end time.Time, duration time.Duration) {
	date := start.Format("2006-01-02")
	logEntry := fmt.Sprintf("- Start: %s | End: %s | Duration: %s\n",
		start.Format("15:04:05"),
		end.Format("15:04:05"),
		duration.Round(time.Second),
	)

	readmePath, err := findReadme()
	if err != nil {
		fmt.Println("❌ README.md not found in current directory.")
		return
	}

	appendToReadme(readmePath, date, logEntry)
}

func findReadme() (string, error) {
	cwd, _ := os.Getwd()
	files, _ := os.ReadDir(cwd)

	for _, file := range files {
		if strings.ToLower(file.Name()) == "readme.md" {
			return filepath.Join(cwd, file.Name()), nil
		}
	}
	return "", fmt.Errorf("README.md not found")
}

func appendToReadme(path, date, entry string) {
	content, _ := os.ReadFile(path)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("❌ Failed to open README.md")
		return
	}
	defer file.Close()

	if !strings.Contains(string(content), "## "+date) {
		file.WriteString(fmt.Sprintf("\n## %s\n", date))
	}
	file.WriteString(entry)
}

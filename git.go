package git

import (
	"fmt"
	"os"
	"os/exec"
)

// CloneRepo клонирует репозиторий и возвращает путь к директории
func CloneRepo(repoURL, branch string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "repo-")
	if err != nil {
		return "", fmt.Errorf("ошибка создания временной директории: %v", err)
	}

	cmd := exec.Command("git", "clone", "-b", branch, repoURL, tmpDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ошибка клонирования репозитория: %v, вывод: %s", err, string(output))
	}

	return tmpDir, nil
}

// CleanupRepo удаляет временную директорию репозитория
func CleanupRepo(dir string) {
	_ = os.RemoveAll(dir)
}

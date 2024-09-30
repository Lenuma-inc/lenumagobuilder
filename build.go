package build

import (
	"fmt"
	"os/exec"
)

// RunBuild выполняет сборку проекта
func RunBuild(repoDir string, buildCmd []string) error {
	if len(buildCmd) == 0 {
		return fmt.Errorf("команда сборки не указана")
	}

	cmd := exec.Command(buildCmd[0], buildCmd[1:]...)
	cmd.Dir = repoDir
	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("ошибка сборки: %v", err)
	}

	return nil
}

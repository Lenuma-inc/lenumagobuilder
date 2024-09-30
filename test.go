package test

import (
	"fmt"
	"os/exec"
)

// RunTests выполняет тесты проекта
func RunTests(repoDir string, testCmd []string) error {
	if len(testCmd) == 0 {
		return fmt.Errorf("команда тестирования не указана")
	}

	cmd := exec.Command(testCmd[0], testCmd[1:]...)
	cmd.Dir = repoDir
	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("ошибка тестирования: %v", err)
	}

	return nil
}

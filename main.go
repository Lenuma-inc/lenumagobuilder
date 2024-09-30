package main

import (
	"build-system/internal/build"
	"build-system/internal/git"
	"build-system/internal/test"
	"log"
)

func main() {
	repoURL := "https://github.com/user/repository.git"
	branch := "main"

	// Клонирование репозитория
	repoDir, err := git.CloneRepo(repoURL, branch)
	if err != nil {
		log.Fatalf("Ошибка клонирования: %v", err)
	}
	defer git.CleanupRepo(repoDir)

	// Сборка проекта
	if err := build.RunBuild(repoDir, []string{"go", "build", "./..."}); err != nil {
		log.Fatalf("Ошибка сборки: %v", err)
	}

	// Запуск тестов
	if err := test.RunTests(repoDir, []string{"go", "test", "./..."}); err != nil {
		log.Fatalf("Ошибка тестирования: %v", err)
	}

	log.Println("Сборка и тесты завершены успешно.")
}

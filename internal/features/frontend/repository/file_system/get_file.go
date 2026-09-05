package frontend_repository

import (
	"errors"
	"fmt"
	"os"
	"path"

	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

func (r *FrontendRepository) GetFile(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)
	fmt.Println(os.ReadDir(path.Join(os.Getenv("PROJECT_ROOT"), "/frontend")))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("faile not found: %w", core_errors.ErrNotFound)
		}

		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return file, nil
}

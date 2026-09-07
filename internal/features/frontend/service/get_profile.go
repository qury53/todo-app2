package frontend_service

import (
	"os"
	"path"
)

func (s *FrontendService) GetProfile() ([]byte, error) {
	filePath := path.Join(os.Getenv("PROJECT_ROOT"), "frontend/profile.html")

	file, err := s.FrontendRepository.GetFile(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

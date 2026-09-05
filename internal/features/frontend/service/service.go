package frontend_service

type FrontendService struct {
	FrontendRepository FrontendRepository
}

type FrontendRepository interface {
	GetFile(fileName string) ([]byte, error)
}

func NewFrontendService(repo FrontendRepository) *FrontendService {
	return &FrontendService{
		FrontendRepository: repo,
	}
}

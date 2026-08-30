package users_service

import "context"

func (s *UserService) DeleteUser(ctx context.Context, userID int) error {
	if err := s.UserRepository.DeleteUser(ctx, userID); err != nil {
		return err
	}

	return nil
}

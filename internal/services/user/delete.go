package services

func (s *UserService) Delete(id string) error {
	if _, err := s.UserRepository.GetById(id); err != nil {
		return err
	}

	if err := s.UserRepository.Delete(id); err != nil {
		return err
	}

	return nil
}

package services

func (s *TaskService) Delete(id string) error {
	if _, err := s.TaskRepository.GetById(id); err != nil {
		return err
	}
	
	if err := s.TaskRepository.Delete(id); err != nil {
		return err
	}

	return nil
}

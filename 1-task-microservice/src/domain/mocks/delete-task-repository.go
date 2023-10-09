package mocks

type deleteTaskRepository struct{}

func NewDeleteTaskRepository() deleteTaskRepository {
	return deleteTaskRepository{}
}

func (r deleteTaskRepository) Execute(id uint) error {
	return nil
}

package contracts

type DeleteTaskRepository interface {
	Execute(id uint) error
}

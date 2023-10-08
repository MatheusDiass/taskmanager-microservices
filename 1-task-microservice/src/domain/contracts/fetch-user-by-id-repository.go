package contracts

type FetchUserByIdRepository interface {
	Execute(id uint) (uint, error)
}

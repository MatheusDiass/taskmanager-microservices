package mocks

type FetchUserByIdRepositoryMock struct {
}

func NewFetchUserByIdRepositoryMock() *FetchUserByIdRepositoryMock {
	return &FetchUserByIdRepositoryMock{}
}

func (r FetchUserByIdRepositoryMock) Execute(id uint) (uint, error) {
	if id == 1 {
		return id, nil
	}

	return 0, nil
}

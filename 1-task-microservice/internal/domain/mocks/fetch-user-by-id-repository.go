package mocks

type fetchUserByIdRepositoryMock struct {
}

func NewFetchUserByIdRepositoryMock() fetchUserByIdRepositoryMock {
	return fetchUserByIdRepositoryMock{}
}

func (r fetchUserByIdRepositoryMock) Execute(id uint) (uint, error) {
	if id == 1 {
		return id, nil
	}

	return 0, nil
}

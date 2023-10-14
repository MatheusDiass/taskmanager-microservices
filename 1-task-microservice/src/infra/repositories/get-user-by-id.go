package repositories

import (
	"task-microservice/src/domain/dtos"
	"task-microservice/src/infra/adapters"
)

type GetUserByIdRepository struct {
	databaseAdapter adapters.GormAdapter
}

func NewGetUserByIdRepository(
	databaseAdapter adapters.GormAdapter,
) GetUserByIdRepository {
	return GetUserByIdRepository{databaseAdapter}
}

func (repository GetUserByIdRepository) Execute(id uint) (uint, error) {
	var user dtos.User
	result := repository.databaseAdapter.GetDb().Table("sch_user.tbl_user").Select("id").First(&user, id)

	if result.Error != nil {
		return 0, result.Error
	}

	return user.Id, nil
}

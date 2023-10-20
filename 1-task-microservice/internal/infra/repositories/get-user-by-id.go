package repositories

import (
	"github.com/matheusdiass/taskmanager-microservices/internal/domain/dtos"
	"github.com/matheusdiass/taskmanager-microservices/internal/infra/adapters"
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
		return 0, nil
	}

	return user.Id, nil
}

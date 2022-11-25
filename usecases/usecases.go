package usecases

import (
	"privy-test/entities"
	"privy-test/repositories"
)

type IUsecases struct {
	repo repositories.RepositoriesInterface
}

type IUsecasesInterface interface {
	CakeUsecases
}

func NewUsecase(repo repositories.RepositoriesInterface) IUsecasesInterface {
	return &IUsecases{
		repo: repo,
	}
}

type CakeUsecases interface {
	CreateCake(req entities.CakeRequest) error
	GetAllCakes(sorting string) (res []entities.CakeResponse, err error)
	GetCakeByID(ID int) (res entities.CakeResponse, err error)
	UpdateCakeByID(ID int, req entities.CakeRequest) error
	DeleteCakeByID(ID int) error
}

package usecases

import (
	"privy-test/entities"
)

func (c IUsecases) CreateCake(req entities.CakeRequest) error {
	err := c.repo.CreateCake(req)
	if err != nil {
		return err
	}

	return nil
}

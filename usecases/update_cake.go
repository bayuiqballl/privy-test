package usecases

import (
	"privy-test/entities"
	"privy-test/presenters"
)

func (c IUsecases) UpdateCakeByID(ID int, req entities.CakeRequest) error {
	res, err := c.repo.GetCakeByID(ID)
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return presenters.ErrIdNotFound
	}

	err = c.repo.UpdateCakeByID(ID, req)
	if err != nil {
		return err
	}

	return nil

}

package usecases

import (
	"fmt"
	"privy-test/presenters"
)

func (c IUsecases) DeleteCakeByID(ID int) error {
	res, err := c.repo.GetCakeByID(ID)
	if err != nil {
		return err
	}

	if res.ID == 0 {
		return presenters.ErrIdNotFound
	}
	fmt.Println(ID)
	err = c.repo.DeleteCakeByID(ID)
	if err != nil {
		return err
	}
	return nil
}

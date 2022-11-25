package usecases

import "privy-test/entities"

func (c IUsecases) GetCakeByID(ID int) (res entities.CakeResponse, err error) {
	res, err = c.repo.GetCakeByID(ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

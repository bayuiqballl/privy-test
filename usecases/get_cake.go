package usecases

import "privy-test/entities"

func (c IUsecases) GetAllCakes(sorting string) (res []entities.CakeResponse, err error) {
	res, err = c.repo.GetAllCakes(sorting)
	if err != nil {
		return res, err
	}

	return res, nil
}

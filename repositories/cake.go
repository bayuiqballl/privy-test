package repositories

import (
	"privy-test/entities"
	query "privy-test/repositories/query"
)

type CakeRepository interface {
	CreateCake(req entities.CakeRequest) error
	GetAllCakes(sorting string) (res []entities.CakeResponse, err error)
	GetCakeByID(ID int) (res entities.CakeResponse, err error)
	UpdateCakeByID(ID int, req entities.CakeRequest) error
	DeleteCakeByID(ID int) error
}

func (c *Repository) CreateCake(req entities.CakeRequest) error {
	_, err := c.db.Exec(
		query.InsertDataCake,
		req.Title,
		req.Description,
		req.Rating,
		req.Image,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *Repository) GetAllCakes(sorting string) (res []entities.CakeResponse, err error) {
	var qry = query.GetAllDataCakes
	if sorting != "" {
		qry += "ORDER BY " + sorting + " ASC"
	}

	row, err := c.db.Query(
		qry,
	)
	if err != nil {
		return res, err
	}

	for row.Next() {
		temp := entities.CakeResponse{}
		err = row.Scan(
			&temp.ID,
			&temp.Title,
			&temp.Description,
			&temp.Rating,
			&temp.Image,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}
	return res, nil
}

func (c *Repository) GetCakeByID(ID int) (res entities.CakeResponse, err error) {
	err = c.db.QueryRow(
		query.GetCakeDetailByID,
		ID,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Description,
		&res.Rating,
		&res.Image,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *Repository) UpdateCakeByID(ID int, req entities.CakeRequest) error {

	_, err := c.db.Exec(
		query.UpdateCakeByID,
		req.Title,
		req.Description,
		req.Rating,
		req.Image,
		ID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (c *Repository) DeleteCakeByID(ID int) error {

	_, err := c.db.Exec(
		query.DeleteCakeByID,
		ID,
	)

	if err != nil {
		return err
	}
	return nil
}

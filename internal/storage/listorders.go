package storage

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func (stor Storage) listOrder() ([]models.DataUnit, error) {
	res := make([]models.DataUnit, 0)
	var temp OrderRecord
	rows, err := stor.db.Query("SELECT * FROM order_data ")
	if err != nil {
		return nil, errors.OpenErr
	}
	defer rows.Close()
	for rows.Next() {
		errScan := rows.Scan(&temp.IdOrder, &temp.IdUser, &temp.IdPackage, &temp.DeliveredDate, &temp.ReceivedDate, &temp.DeadLine, &temp.RefundDate, &temp.Mass)
		if errScan != nil {
			return nil, errors.NotResolverErr
		}
		res = append(res, temp.ToDomain())
	}
	return res, nil
}

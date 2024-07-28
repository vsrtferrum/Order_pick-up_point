package storage

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func (stor Storage) ChangePackge(data models.ChangePackage) error {
	val, flag := stor.cacheOrder.Get(data.IdOrder)
	if !flag {
		return errors.FoundOrderEr
	}

	val.IdPackage = data.IdPackage
	stor.cacheOrder.Set(data.IdOrder, val)
	tx, err := stor.db.Begin()
	if err != nil {
		return err
	}
	_, errQ := tx.Exec("UPDATE order_data SET id_package = $1 WHERE id_order = $2;",
		data.IdPackage, data.IdOrder)
	if errQ != nil {
		tx.Rollback()
	}
	errC := tx.Commit()
	return errC
}

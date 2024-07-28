package storage

import "gitlab.ozon.dev/berkinv/homework/internal/errors"

func (stor Storage) DeleteRowOrderData(ido uint32) error {
	stor.cacheOrder.Delete(ido)
	tx, err := stor.db.Begin()
	if err != nil {
		return err
	}
	_, errQ := tx.Exec("DELETE FROM order_data WHERE id_order = $1", ido)
	if errQ != nil {
		tx.Rollback()
		return errors.CantResolvArgsErr
	}
	errC := tx.Commit()
	if errC != nil {
		return errors.CantResolvArgsErr
	}
	return nil
}

package storage

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func (stor Storage) UpdateOrderData(unit models.DataUnit) error {
	tx, err := stor.db.Begin()
	if err != nil {
		return err
	}
	_, errQ := tx.Exec("UPDATE order_data SET delivered_date = $1, recieved_date =$2,  dead_line = $3, refund_date =$4 WHERE id_order = $5;",
		unit.DeliveredDate, unit.ReceivedDate, unit.DeadLine, unit.RefundDate, unit.IdOrder)
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

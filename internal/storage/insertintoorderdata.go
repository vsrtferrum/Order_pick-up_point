package storage

import "gitlab.ozon.dev/berkinv/homework/internal/models"

func (stor Storage) InsertIntoOrderData(d models.DataUnit) error {
	stor.cacheOrder.Set(d.IdOrder, d)
	tx, err := stor.db.Begin()
	if err != nil {
		return err
	}
	temp := Transofrm(d)
	_, errQ := tx.Exec("INSERT INTO order_data(id_order, id_user, id_package, delivered_date, recieved_date, dead_line, refund_date, item_mass)values($1, $2, $3, $4, $5, $6, $7, $8);",
		temp.IdOrder, temp.IdUser, temp.IdPackage, temp.DeliveredDate, temp.ReceivedDate, temp.DeadLine, temp.RefundDate, temp.Mass)
	if errQ != nil {
		tx.Rollback()
		return errQ
	}
	err = tx.Commit()
	return err
}

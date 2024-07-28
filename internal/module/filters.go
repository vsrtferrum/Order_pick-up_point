package module

import (
	"time"

	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

type filters interface {
	CheckRefundOrderDeliver(res []models.DataUnit, data models.ReceiveOrderDeliver) (models.DataUnit, error)
	CheckOrderList(mas []models.DataUnit, data models.OrderList) []models.DataUnit
}

func CheckRefundOrderDeliver(res []models.DataUnit, data models.ReceiveOrderDeliver) (models.DataUnit, error) {
	if len(res) == 0 {
		return models.DataUnit{}, errors.FoundOrderEr
	} else if len(res) == 1 {
		if res[0].IdOrder == data.IdOrder && checkrefund(res[0]) {
			res[0].ReceivedDate = time.Now()
			res[0].DeadLine = time.Now().AddDate(0, 0, data.DeadLine)
			return res[0], nil
		} else {
			return models.DataUnit{}, errors.CantResolvArgsErr
		}
	} else {
		return models.DataUnit{}, errors.CantResolvArgsErr
	}
}

func CheckOrderList(mas []models.DataUnit, data models.OrderList) []models.DataUnit {
	index := 0
	ans := make([]models.DataUnit, len(mas))
	for _, val := range mas {
		if val.IdUser == data.IdUser && val.DeadLine.After(time.Now()) &&
			val.RefundDate.IsZero() && val.DeliveredDate.Before(time.Now()) {
			ans[index] = val
			index++
		}
	}
	return ans[:index]
}

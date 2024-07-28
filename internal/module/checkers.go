package module

import (
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func checkrefund(val models.DataUnit) bool {
	if val.DeliveredDate.IsZero() && val.DeadLine.IsZero() &&
		val.RefundDate.IsZero() && val.ReceivedDate.IsZero() {
		return true
	}
	return false
}

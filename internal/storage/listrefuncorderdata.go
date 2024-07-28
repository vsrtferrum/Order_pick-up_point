package storage

import (
	"time"

	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func (stor Storage) ListRefuncOrderData(equal time.Time) ([]models.DataUnit, error) {
	cnt := 0
	data := stor.cacheOrder.GetAll()
	ans := make([]models.DataUnit, len(data))
	for _, v := range data {
		if !v.RefundDate.IsZero() {
			ans[cnt] = v
			cnt++
		}
	}
	return ans[:cnt], nil
}

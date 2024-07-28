package storage

import (
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func (stor Storage) ListUsersOrderData(equal uint32) ([]models.DataUnit, error) {
	cnt := 0
	all := stor.cacheOrder.GetAll()
	result := make([]models.DataUnit, len(all))
	for _, v := range all {
		if v.IdUser == equal {
			result[cnt] = v
			cnt++
		}
	}
	return result[:cnt], nil
}

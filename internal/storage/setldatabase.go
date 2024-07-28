package storage

import (
	"gitlab.ozon.dev/berkinv/homework/internal/cache"
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func (stor *Storage) SetFullDatabaseReq(req string, limit int) error {
	stor.fullDatabaseReq = req
	var err error
	stor.db, err = Open(stor.fullDatabaseReq)
	if err != nil {
		return errors.OpenErr
	}
	stor.cachePackage = cache.NewTTLClient[uint32, models.PackageUnit](limit)
	stor.cacheOrder = cache.NewTTLClient[uint32, models.DataUnit](limit)
	ans, err := stor.listPackage()
	if err != nil {
		return err
	}
	for _, val := range ans {
		stor.cachePackage.Set(val.IdPackage, val)
	}
	orders, err := stor.listOrder()
	if err != nil {
		return err
	}
	for _, val := range orders {
		stor.cacheOrder.Set(val.IdOrder, val)
	}

	return nil
}

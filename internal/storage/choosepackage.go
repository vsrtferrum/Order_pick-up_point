package storage

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func (stor Storage) ChoosePackage(id_pack uint32) (models.PackageUnit, error) {
	val, err := stor.cachePackage.Get(id_pack)
	if !err {
		return val, errors.CantResolvArgsErr
	}
	return val, nil
}

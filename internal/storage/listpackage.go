package storage

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func (stor Storage) ListPackage() ([]models.PackageUnit, error) {
	res := stor.cachePackage.GetAll()
	ans := make([]models.PackageUnit, len(res))
	cnt := 0
	for _, v := range res {
		ans[cnt] = v
		cnt++
	}
	return ans, nil
}

func (stor Storage) listPackage() ([]models.PackageUnit, error) {
	res := make([]models.PackageUnit, 0)
	var temp packageRecord
	rows, err := stor.db.Query("SELECT * FROM packages;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		errScan := rows.Scan(&temp.IdPackage, &temp.PackageCost, &temp.PackageName, &temp.LowerMass, &temp.UpperMass)
		if errScan != nil {
			return nil, errors.NotResolverErr
		}
		res = append(res, temp.ToDomain())
	}
	return res, nil
}

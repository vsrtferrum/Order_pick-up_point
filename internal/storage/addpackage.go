package storage

import "gitlab.ozon.dev/berkinv/homework/internal/models"

func (stor Storage) AddPackage(d models.PackageUnit) error {
	tx, err := stor.db.Begin()
	if err != nil {
		return err
	}
	temp := TransofrmPack(d)
	_, errQ := tx.Exec("INSERT INTO packages (package_cost, package_name, lower_mass, upper_mass) values ($1, $2, $3, $4);",
		temp.PackageCost, temp.PackageName, temp.LowerMass, temp.UpperMass)
	if errQ != nil {
		tx.Rollback()
		return errQ
	}
	err = tx.Commit()
	stor.cachePackage.Set(d.IdPackage, d)
	return err
}

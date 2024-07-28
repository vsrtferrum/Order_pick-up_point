package storage

import "gitlab.ozon.dev/berkinv/homework/internal/models"

func (t OrderRecord) ToDomain() models.DataUnit {
	return models.DataUnit{
		IdOrder:       t.IdOrder,
		IdUser:        t.IdUser,
		IdPackage:     t.IdPackage,
		DeliveredDate: t.DeliveredDate,
		ReceivedDate:  t.ReceivedDate,
		DeadLine:      t.DeadLine,
		RefundDate:    t.RefundDate,
		Mass:          t.Mass,
	}
}
func Transofrm(models models.DataUnit) *OrderRecord {
	return &OrderRecord{
		IdOrder:       models.IdOrder,
		IdUser:        models.IdUser,
		IdPackage:     models.IdPackage,
		DeliveredDate: models.DeliveredDate,
		ReceivedDate:  models.ReceivedDate,
		DeadLine:      models.DeadLine,
		RefundDate:    models.RefundDate,
		Mass:          models.Mass,
	}
}
func (t packageRecord) ToDomain() models.PackageUnit {
	return models.PackageUnit{
		PackageCost: t.PackageCost,
		PackageName: t.PackageName,
		LowerMass:   t.LowerMass,
		UpperMass:   t.UpperMass,
	}
}
func TransofrmPack(models models.PackageUnit) *packageRecord {
	return &packageRecord{
		PackageCost: models.PackageCost,
		PackageName: models.PackageName,
		LowerMass:   models.LowerMass,
		UpperMass:   models.UpperMass,
	}
}

package storage

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"gitlab.ozon.dev/berkinv/homework/internal/cache"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

type ReadnWrite interface {
	SetFullDatabaseReq()
	InsertIntoOrderData(d models.DataUnit) error
	UpdateOrderData(unit models.DataUnit) error
	ListUsersOrderData(equal uint) ([]models.DataUnit, error)
	ListRefuncOrderData(equal time.Time) ([]models.DataUnit, error)
	ListOrderData(equal uint) ([]models.DataUnit, error)
	DeleteRowOrderData(ido uint) error
	ListPackage() ([]models.PackageUnit, error)
	ListOrders() ([]models.PackageUnit, error)
	ChoosePackage(id_pack uint) (models.PackageUnit, error)
	ChangePackge(data models.ChangePackage) error
	AddPackage(data models.PackageUnit) error
	GetDb() *sql.DB
}

type Storage struct {
	fullDatabaseReq string
	db              *sql.DB
	cacheOrder      *cache.TTLClient[uint32, models.DataUnit]
	cachePackage    *cache.TTLClient[uint32, models.PackageUnit]
	ReadnWrite
}

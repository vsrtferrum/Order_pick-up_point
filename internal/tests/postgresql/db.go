package postgresql

import (
	"gitlab.ozon.dev/berkinv/homework/internal/storage"
	"testing"
)

type TDB struct {
	DB storage.Storage
}

func NewFromEnv() *TDB {
	db := storage.Storage{}
	err := db.SetFullDatabaseReq("user=vsrtf dbname=postgres  sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &TDB{DB: db}
}
func (d *TDB) SetUp(t *testing.T) {
	t.Helper()
}
func (d *TDB) TearDown(t *testing.T) {
	t.Helper()
}
func (d *TDB) TruncateTables() error {

	tx, err := d.DB.GetDb().Begin()
	if err != nil {
		return err
	}
	if _, errD := tx.Exec("TRUNCATE order_data, packages RESTART IDENTITY"); errD != nil {
		tx.Rollback()
		return errD
	}
	errC := tx.Commit()
	return errC
}

package storage

import "database/sql"

func (stor *Storage) GetDb() *sql.DB {
	return stor.db
}

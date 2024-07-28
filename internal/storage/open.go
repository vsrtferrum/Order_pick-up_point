package storage

import "database/sql"

func Open(req string) (*sql.DB, error) {
	db, err := sql.Open("postgres", req)
	if err != nil {
		return nil, err
	}
	return db, nil
}

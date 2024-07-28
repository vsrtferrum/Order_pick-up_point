//go:build integration

package tests

import "gitlab.ozon.dev/berkinv/homework/internal/tests/postgresql"

var (
	db *postgresql.TDB
)

func init() {
	db = postgresql.NewFromEnv()
}

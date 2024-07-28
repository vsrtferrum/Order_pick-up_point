//go:build integration

package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
	"gitlab.ozon.dev/berkinv/homework/internal/tests/postgresql"
	"testing"
	"time"
)

var (
	zeroTime time.Time
	packs    = make([]models.PackageUnit, 3)
)

func Clear(t *testing.T) {
	db.TruncateTables()
}
func TestClear(t *testing.T) {
	db.TruncateTables()
}
func PackageData(t *testing.T) {
	//arr
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	//act
	errtape := repo.DB.AddPackage(
		models.PackageUnit{
			PackageCost: 1,
			PackageName: "tape",
			LowerMass:   0,
			UpperMass:   32766,
		})
	errbox := repo.DB.AddPackage(
		models.PackageUnit{
			PackageCost: 20,
			PackageName: "box",
			LowerMass:   0,
			UpperMass:   30,
		})
	errpack := repo.DB.AddPackage(
		models.PackageUnit{
			PackageCost: 5,
			PackageName: "box",
			LowerMass:   0,
			UpperMass:   10,
		})
	//assert
	require.NoError(t, errtape)
	require.NoError(t, errbox)
	require.NoError(t, errpack)
}
func OrderData(t *testing.T) {
	//arr
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	//act
	err1 := repo.DB.InsertIntoOrderData(
		models.DataUnit{
			IdOrder:       1,
			IdUser:        1,
			IdPackage:     1,
			ReceivedDate:  zeroTime,
			DeliveredDate: zeroTime,
			DeadLine:      time.Now().Add(2 * time.Minute),
			RefundDate:    zeroTime,
			Mass:          100,
		})
	err2 := repo.DB.InsertIntoOrderData(
		models.DataUnit{
			IdOrder:       2,
			IdUser:        1,
			IdPackage:     1,
			ReceivedDate:  zeroTime,
			DeliveredDate: zeroTime,
			DeadLine:      time.Now().Add(2 * time.Hour),
			RefundDate:    zeroTime,
			Mass:          9,
		})
	err3 := repo.DB.InsertIntoOrderData(
		models.DataUnit{
			IdOrder:       3,
			IdUser:        1,
			IdPackage:     1,
			ReceivedDate:  zeroTime,
			DeliveredDate: zeroTime,
			DeadLine:      time.Now().Add(2 * time.Second),
			RefundDate:    zeroTime,
			Mass:          5,
		})
	err4 := repo.DB.InsertIntoOrderData(
		models.DataUnit{
			IdOrder:       4,
			IdUser:        2,
			IdPackage:     1,
			ReceivedDate:  zeroTime,
			DeliveredDate: zeroTime,
			DeadLine:      time.Now().Add(2 * time.Hour),
			RefundDate:    zeroTime,
			Mass:          25,
		})
	//assert
	require.NoError(t, err1)
	require.NoError(t, err2)
	require.NoError(t, err3)
	require.NoError(t, err4)
}
func RefundData(t *testing.T) {
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	repo.DB.UpdateOrderData(models.DataUnit{
		IdOrder:       3,
		IdUser:        1,
		IdPackage:     1,
		ReceivedDate:  time.Now(),
		DeliveredDate: time.Now(),
		DeadLine:      time.Now().Add(2 * time.Second),
		RefundDate:    time.Now(),
		Mass:          5,
	})
}
func TestInsertOrderData(t *testing.T) {
	PackageData(t)
	OrderData(t)
	t.Parallel()
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	err1 := repo.DB.InsertIntoOrderData(models.DataUnit{
		IdOrder:       5,
		IdUser:        1,
		IdPackage:     2,
		ReceivedDate:  time.Now(),
		DeliveredDate: time.Now(),
		DeadLine:      time.Now(),
		RefundDate:    zeroTime,
		Mass:          5,
	})
	err2 := repo.DB.InsertIntoOrderData(models.DataUnit{
		IdOrder:       1,
		IdUser:        1,
		IdPackage:     2,
		ReceivedDate:  time.Now(),
		DeliveredDate: time.Now(),
		DeadLine:      time.Now(),
		RefundDate:    zeroTime,
		Mass:          5,
	})
	require.NoError(t, err1)
	require.Error(t, err2)
	Clear(t)
}
func TestUpdateOrderData(t *testing.T) {
	PackageData(t)
	OrderData(t)
	t.Parallel()
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	err1 := repo.DB.UpdateOrderData(models.DataUnit{
		IdOrder:       3,
		IdUser:        1,
		IdPackage:     2,
		ReceivedDate:  time.Now(),
		DeliveredDate: time.Now(),
		DeadLine:      time.Now(),
		RefundDate:    zeroTime,
		Mass:          5,
	})
	err2 := repo.DB.UpdateOrderData(models.DataUnit{
		IdOrder:       1000,
		IdUser:        1,
		IdPackage:     2,
		ReceivedDate:  time.Now(),
		DeliveredDate: time.Now(),
		DeadLine:      time.Now(),
		RefundDate:    zeroTime,
		Mass:          5,
	})
	require.NoError(t, err1)
	require.NoError(t, err2)
	Clear(t)

}
func TestListRefund(t *testing.T) {
	PackageData(t)
	OrderData(t)
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	ans, err := repo.DB.ListRefuncOrderData(zeroTime)
	require.NoError(t, err)
	assert.Len(t, ans, 0)
	RefundData(t)
	ans2, err2 := repo.DB.ListRefuncOrderData(zeroTime)
	require.NoError(t, err2)
	assert.Len(t, ans2, 1)
	assert.Equal(t, uint(3), ans2[0].IdOrder)
	Clear(t)
}
func TestChange(t *testing.T) {
	PackageData(t)
	OrderData(t)
	t.Parallel()
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	err := repo.DB.ChangePackge(
		models.ChangePackage{
			IdOrder:   3,
			IdPackage: 2,
		})
	require.NoError(t, err)
	Clear(t)
}
func TestChoose(t *testing.T) {
	PackageData(t)
	OrderData(t)
	t.Parallel()
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	ans, err := repo.DB.ChoosePackage(1)
	_, err2 := repo.DB.ChoosePackage(4)
	require.NoError(t, err)
	require.Equal(t, errors.NotResolverErr, err2)
	assert.Equal(t, ans, models.PackageUnit{
		PackageCost: 1,
		PackageName: "tape",
		LowerMass:   0,
		UpperMass:   32766,
	})
	Clear(t)
}
func TestListOrder(t *testing.T) {
	PackageData(t)
	OrderData(t)
	t.Parallel()
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	ans, err := repo.DB.ListOrderData(1)
	ans2, err2 := repo.DB.ListOrderData(1080)
	require.NoError(t, err)
	require.NoError(t, err2)
	assert.Len(t, ans, 1)
	assert.Len(t, ans2, 0)
	Clear(t)
}
func TestListPackage(t *testing.T) {
	PackageData(t)
	OrderData(t)
	t.Parallel()
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	packs[0] = models.PackageUnit{
		PackageCost: 1,
		PackageName: "tape",
		LowerMass:   0,
		UpperMass:   32766}
	packs[1] = models.PackageUnit{
		PackageCost: 20,
		PackageName: "box",
		LowerMass:   0,
		UpperMass:   30}
	packs[2] = models.PackageUnit{
		PackageCost: 5,
		PackageName: "box",
		LowerMass:   0,
		UpperMass:   10}
	ans, err := repo.DB.ListPackage()
	require.NoError(t, err)
	assert.Equal(t, packs, ans)
	Clear(t)
}
func TestDeleteOrderData(t *testing.T) {
	PackageData(t)
	OrderData(t)
	t.Parallel()
	db.SetUp(t)
	defer db.TearDown(t)
	repo := postgresql.NewFromEnv()
	err := repo.DB.DeleteRowOrderData(3)
	err2 := repo.DB.DeleteRowOrderData(110)
	require.NoError(t, err)
	require.NoError(t, err2)
	Clear(t)
}

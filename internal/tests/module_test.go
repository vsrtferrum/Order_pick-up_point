package tests

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
	"gitlab.ozon.dev/berkinv/homework/internal/module"
	mock_module "gitlab.ozon.dev/berkinv/homework/internal/module/mocks"
)

var (
	zeroTime time.Time
	packs    = make([]models.DataUnit, 3)
)

func TestCheckOL(t *testing.T) {
	packs[0] = models.DataUnit{
		IdOrder:       1,
		IdUser:        1,
		IdPackage:     1,
		ReceivedDate:  zeroTime,
		DeliveredDate: time.Now(),
		DeadLine:      time.Now().Add(2 * time.Minute),
		RefundDate:    zeroTime,
		Mass:          100,
	}
	packs[1] = models.DataUnit{
		IdOrder:       2,
		IdUser:        1,
		IdPackage:     1,
		ReceivedDate:  zeroTime,
		DeliveredDate: time.Now(),
		DeadLine:      time.Now().Add(2 * time.Hour),
		RefundDate:    zeroTime,
		Mass:          9,
	}
	packs[2] = models.DataUnit{
		IdOrder:       3,
		IdUser:        1,
		IdPackage:     1,
		ReceivedDate:  zeroTime,
		DeliveredDate: time.Now(),
		DeadLine:      time.Now().Add(2 * time.Second),
		RefundDate:    zeroTime,
		Mass:          5,
	}
	t.Run("Time Ended test", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		mockMod := mock_module.NewMockfilters(ctrl)
		mockMod.EXPECT().CheckOrderList(gomock.Any(), gomock.Any()).AnyTimes()

		time.Sleep(2 * time.Second)
		res1 := module.CheckOrderList(packs, models.OrderList{IdUser: 1})
		//res2 := mockMod.CheckOL(packs, models.OrderList{IdUser: 100})
		assert.Equal(t, packs[:2], res1)
		//assert.Equal(t, ans_res2, res2)

	})
	t.Run("Empty check test", func(t *testing.T) {
		t.Parallel()
		ans := make([]models.DataUnit, 0)
		// arrange
		ctrl := gomock.NewController(t)
		mockMod := mock_module.NewMockfilters(ctrl)
		mockMod.EXPECT().CheckOrderList(gomock.Any(), gomock.Any()).AnyTimes()
		res1 := module.CheckOrderList(packs, models.OrderList{IdUser: 10})
		assert.Equal(t, ans, res1)
	})
}

func TestCheckROD(t *testing.T) {
	packs[0] = models.DataUnit{
		IdOrder:       1,
		IdUser:        1,
		IdPackage:     1,
		ReceivedDate:  zeroTime,
		DeliveredDate: zeroTime,
		DeadLine:      zeroTime,
		RefundDate:    zeroTime,
		Mass:          100,
	}
	packs[1] = models.DataUnit{
		IdOrder:       2,
		IdUser:        1,
		IdPackage:     1,
		ReceivedDate:  time.Now(),
		DeliveredDate: zeroTime,
		DeadLine:      zeroTime,
		RefundDate:    zeroTime,
		Mass:          9,
	}
	packs[2] = models.DataUnit{
		IdOrder:       3,
		IdUser:        1,
		IdPackage:     1,
		ReceivedDate:  zeroTime,
		DeliveredDate: zeroTime,
		DeadLine:      zeroTime,
		RefundDate:    zeroTime,
		Mass:          5,
	}
	t.Run("Test with edit", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		mockMod := mock_module.NewMockfilters(ctrl)
		mockMod.EXPECT().CheckRefundOrderDeliver(gomock.Any(), gomock.Any()).AnyTimes()
		ans, err := module.CheckRefundOrderDeliver(packs, models.ReceiveOrderDeliver{IdOrder: 3, IdUser: 1, IdPackage: 1, DeadLine: 2, Mass: 5})
		assert.NoError(t, err)
		require.Equal(t, uint32(3), ans.IdOrder)

	})
	t.Run("Test with insert", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		mockMod := mock_module.NewMockfilters(ctrl)
		mockMod.EXPECT().CheckRefundOrderDeliver(gomock.Any(), gomock.Any()).AnyTimes()
		ans, err := module.CheckRefundOrderDeliver(packs, models.ReceiveOrderDeliver{IdOrder: 5, IdUser: 1, IdPackage: 1, DeadLine: 2, Mass: 5})
		assert.NoError(t, err)
		require.Equal(t, uint32(0), ans.IdOrder)

	})
	t.Run("Test with non empty time", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		mockMod := mock_module.NewMockfilters(ctrl)
		mockMod.EXPECT().CheckRefundOrderDeliver(gomock.Any(), gomock.Any()).AnyTimes()
		_, err := module.CheckRefundOrderDeliver(packs, models.ReceiveOrderDeliver{IdOrder: 2, IdUser: 1, IdPackage: 1, DeadLine: 2, Mass: 5})
		assert.Equal(t, errors.CantResolvArgsErr, err)

	})
}

package tests

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pvz "gitlab.ozon.dev/berkinv/homework/pkg/api/proto/pvz/v1/pvz/v1"

	mock_service "gitlab.ozon.dev/berkinv/homework/internal/api/mock"
)

func TestAddPackage(t *testing.T) {
	t.Run("smoke test", func(t *testing.T) {
		t.Parallel()
		req := &pvz.AddPackageRequest{
			PackageName: "test",
			LowerMass:   uint32(0),
			UpperMass:   uint32(10),
			PackageCost: uint32(1),
		}
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		pvz := mock_service.NewMockImplementation(ctrl)
		pvz.EXPECT().AddPackage(gomock.Any(), gomock.Any()).AnyTimes()
		res, err := pvz.AddPackage(ctx, req)
		assert.NoError(t, err)
		require.Nil(t, res)

	})
}
func TestOrderlis(t *testing.T) {
	t.Run("smoke test", func(t *testing.T) {
		t.Parallel()
		req := &pvz.OrderListRequest{
			IdUser: uint32(1),
		}
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		pvz := mock_service.NewMockImplementation(ctrl)
		pvz.EXPECT().OrderList(gomock.Any(), gomock.Any()).AnyTimes()
		res, err := pvz.OrderList(ctx, req)
		assert.NoError(t, err)
		require.Nil(t, res)

	})
}

package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/berkinv/homework/internal/api"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
	"gitlab.ozon.dev/berkinv/homework/pkg/api/proto/pvz/v1/pvz/v1"
)

func TestPackageToDomain(t *testing.T) {
	t.Run("Test pkg converter", func(t *testing.T) {
		t.Parallel()
		data := pvz.AddPackageRequest{
			PackageName: "test",
			PackageCost: 1,
			LowerMass:   0,
			UpperMass:   10,
		}
		res := api.PackageToDomain(&data)
		require.Equal(t, models.PackageUnit{
			IdPackage:   0,
			PackageName: "test",
			PackageCost: 1,
			LowerMass:   0,
			UpperMass:   10,
		}, res)
	})
}
func TestRefundUserToDomain(t *testing.T) {
	t.Run("Test data unit converter", func(t *testing.T) {
		t.Parallel()
		res := api.RefundUserToDomain(&pvz.RefundUserRequest{IdUser: 1, IdOrder: 1})
		assert.Equal(t, models.RefundUser{
			IdOrder: 1,
			IdUser:  1,
		}, res)

	})
}

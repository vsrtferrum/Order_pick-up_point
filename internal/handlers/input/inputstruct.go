package input

import (
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

type reader interface {
	CliChangePack(arg []string) (models.ChangePackage, error)
	CliAccept(arg []string) (models.ReceiveOrderDeliver, error)
	CliRefundDeliver(arg []string) (models.RefundOrderDeliver, error)
	CliIssueUser(arg []string) ([]models.ReceiveOrderUser, error)
	CliListOrder(arg []string) (models.OrderList, uint64, error)
	CliRefundUser(arg []string) (models.RefundUser, error)
	CliListRefund(arg []string) (uint64, error)
	CliSetWorkersNum(arg []string) (uint64, error)
	CliAddPackage(arg []string) (models.PackageUnit, error)
}

type Input struct {
	reader
	sw bool
}

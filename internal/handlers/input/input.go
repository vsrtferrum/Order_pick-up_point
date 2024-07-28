package input

import (
	"fmt"
	"strconv"

	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func CliChangePack(arg []string) (models.ChangePackage, error) {
	if len(arg) != 2 {
		return models.ChangePackage{}, errors.CantResolvArgsErr
	}
	t1, err1 := strconv.ParseUint(arg[0], 10, 16)
	t2, err2 := strconv.ParseUint(arg[1], 10, 16)
	if err1 != nil || err2 != nil {
		return models.ChangePackage{}, errors.CantResolvArgsErr
	}

	return models.ChangePackage{IdOrder: uint32(t1), IdPackage: uint32(t2)}, nil
}
func CliAccept(arg []string) (models.ReceiveOrderDeliver, error) {
	if len(arg) != 4 {
		return models.ReceiveOrderDeliver{}, errors.CantResolvArgsErr
	}
	t1, err1 := strconv.ParseUint(arg[0], 10, 16)
	t2, err2 := strconv.ParseUint(arg[1], 10, 16)
	t3, err3 := strconv.ParseInt(arg[2], 10, 16)
	t4, err4 := strconv.ParseUint(arg[3], 10, 16)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return models.ReceiveOrderDeliver{}, errors.CantResolvArgsErr
	}
	return models.ReceiveOrderDeliver{IdOrder: uint32(t1), IdUser: uint32(t2), IdPackage: 1,
		DeadLine: int(t3), Mass: uint32(t4)}, nil
}
func CliRefundDeliver(arg []string) (models.RefundOrderDeliver, error) {
	if len(arg) != 1 {
		return models.RefundOrderDeliver{}, errors.CantResolvArgsErr
	}
	t, err1 := strconv.ParseUint(arg[0], 10, 16)
	if err1 != nil {
		fmt.Println("Ошибка в первом аргументе")
		return models.RefundOrderDeliver{}, errors.CantResolvArgsErr
	}

	return models.RefundOrderDeliver{IdOrder: uint32(t)}, nil
}
func CliIssueUser(arg []string) ([]models.ReceiveOrderUser, error) {
	idO := make([]models.ReceiveOrderUser, len(arg))
	for i := range arg {
		temp, err := strconv.ParseUint(arg[i], 10, 8)
		if err == nil {
			idO[i] = models.ReceiveOrderUser{IdOrder: uint32(temp)}
		} else {
			return nil, errors.CantResolvArgsErr
		}
	}
	return idO, nil
}
func CliListOrder(arg []string) (models.OrderList, uint64, error) {
	if len(arg) == 2 {
		t1, err1 := strconv.ParseUint(arg[0], 10, 16)
		t2, err2 := strconv.ParseUint(arg[1], 10, 16)
		if err1 != nil || err2 != nil {
			return models.OrderList{}, 0, errors.CantResolvArgsErr
		}
		return models.OrderList{IdUser: uint32(t1)}, t2, nil
	} else if len(arg) == 1 {
		t1, err1 := strconv.ParseUint(arg[0], 10, 64)
		if err1 != nil {
			return models.OrderList{}, 0, errors.CantResolvArgsErr
		}
		return models.OrderList{IdUser: uint32(t1)}, 0, nil
	}
	return models.OrderList{}, 0, errors.CantResolvArgsErr
}
func CliRefundUser(arg []string) (models.RefundUser, error) {
	if len(arg) != 2 {
		return models.RefundUser{}, errors.CantResolvArgsErr
	}
	t1, err1 := strconv.ParseUint(arg[0], 10, 16)
	t2, err2 := strconv.ParseUint(arg[1], 10, 16)
	if err1 != nil || err2 != nil {
		return models.RefundUser{}, errors.CantResolvArgsErr
	}
	return models.RefundUser{IdUser: uint32(t1), IdOrder: uint32(t2)}, nil
}
func CliListRefund(arg []string) (uint64, error) {
	if len(arg) != 1 {
		return 0, errors.CantResolvArgsErr
	}
	t, err := strconv.ParseUint(arg[0], 10, 64)
	if err != nil {
		return 0, err
	}
	return t, nil
}
func CliSetWorkersNum(arg []string) (uint64, error) {
	return CliListRefund(arg)
}
func CliAddPackage(arg []string) (models.PackageUnit, error) {
	if len(arg) != 4 {
		return models.PackageUnit{}, errors.CantResolvArgsErr
	}
	t1, err1 := strconv.ParseUint(arg[0], 10, 16)
	t2 := arg[1]
	t3, err2 := strconv.ParseUint(arg[2], 10, 16)
	t4, err3 := strconv.ParseUint(arg[3], 10, 16)

	if err1 != nil || err2 != nil || err3 != nil {
		return models.PackageUnit{}, errors.CantResolvArgsErr
	}
	return models.PackageUnit{PackageCost: uint32(t1), PackageName: t2,
		LowerMass: uint32(t3), UpperMass: uint32(t4)}, nil
}

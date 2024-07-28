package module

import (
	"fmt"
	"sort"
	"time"

	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/models"
	"gitlab.ozon.dev/berkinv/homework/internal/storage"
)

var zeroTime time.Time

type ModuleInterface interface {
	ReceiveOrderDeliver(data models.ReceiveOrderDeliver) error
	RefundDeliver(data []models.RefundOrderDeliver) error
	ReceiveOrderUser(data models.ReceiveOrderUser) error
	OrderList(data models.OrderList) ([]models.DataUnit, error)
	RefundUser(data models.RefundUser) error
	RefundList(data models.RefundList) ([]models.DataUnit, error)
	ChangePackage(data models.ChangePackage) error
}

type Module struct {
	ModuleInterface
	storage.Storage
}

func (mod Module) ReceiveOrderDeliver(data models.ReceiveOrderDeliver) error {
	res, err := mod.ListOrderData(data.IdOrder)
	if err != nil {
		return err
	}
	ans, err := CheckRefundOrderDeliver(res, data)
	if err != nil {
		return err
	}
	if ans.IdOrder == 0 {
		errD := mod.InsertIntoOrderData(models.DataUnit{IdUser: data.IdUser, IdOrder: data.IdOrder,
			IdPackage: data.IdPackage, DeliveredDate: time.Now(),
			ReceivedDate: zeroTime, DeadLine: time.Now().AddDate(0, 0, data.DeadLine),
			RefundDate: zeroTime, Mass: data.Mass})
		return errD
	}
	errD := mod.UpdateOrderData(ans)
	return errD
}

func (mod Module) RefundDeliver(data models.RefundOrderDeliver) error {
	flag := false
	res, err := mod.ListOrderData(data.IdOrder)
	if err != nil {
		return err
	}
	for _, val := range res {
		if val.IdOrder == data.IdOrder && val.DeadLine.Before(time.Now()) &&
			!(val.ReceivedDate.IsZero()) && val.RefundDate.IsZero() &&
			val.DeliveredDate.Before(time.Now()) {
			errDelete := mod.DeleteRowOrderData(data.IdOrder)
			if errDelete != nil {
				return errDelete
			}
		} else if !(val.IdOrder == data.IdOrder && val.DeadLine.Before(time.Now()) &&
			!(val.ReceivedDate.IsZero()) && val.RefundDate.IsZero() && val.DeliveredDate.Before(time.Now())) {
			flag = true
		}
	}
	if flag {
		return errors.CantRefundDeliverErr
	}
	return nil
}

func (mod Module) ReceiveOrderUser(data []models.ReceiveOrderUser) error {
	res := make([]models.DataUnit, 0)
	for _, val := range data {
		temp, err := mod.ListUsersOrderData(val.IdOrder)
		if err != nil {
			return err
		}
		res = append(res, temp...)
	}
	if len(res) == 0 {
		return errors.CantRecieveOrderUserErr
	}
	for i, val := range res {
		if val.DeadLine.After(time.Now()) && val.RefundDate.IsZero() && val.DeliveredDate.Before(time.Now()) {
			res[i].ReceivedDate = time.Now()

		}
	}
	for _, val := range res {
		err := mod.UpdateOrderData(val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mod Module) OrderList(data models.OrderList) ([]models.DataUnit, error) {
	res, err := mod.ListUsersOrderData(data.IdUser)
	if err != nil {
		return nil, err
	}
	mas := CheckOrderList(res, data)
	sort.SliceStable(mas, func(i, j int) bool {
		return mas[i].DeliveredDate.After(mas[j].DeliveredDate)
	})
	return mas, nil
}

func (mod Module) RefundUser(data models.RefundUser) error {
	flag := false
	res, errS := mod.ListUsersOrderData(data.IdUser)
	if errS != nil {
		return errS
	}
	for _, val := range res {
		if data.IdOrder == val.IdOrder && val.IdUser == data.IdUser &&
			!(val.DeadLine.IsZero()) && val.RefundDate.IsZero() &&
			val.DeliveredDate.Before(time.Now()) && !(val.ReceivedDate.IsZero()) {
			val.RefundDate = time.Now()
			err := mod.UpdateOrderData(val)
			if err != nil {
				return err
			}
			flag = true
		}
	}
	if flag {
		return nil
	}
	return errors.RefundUserErr
}

func (mod Module) RefundList() ([]models.DataUnit, error) {
	sel, err := mod.ListRefuncOrderData(zeroTime)
	if err != nil {
		return nil, err
	}
	return sel, nil
}

func (mod Module) ChangePackage(data models.ChangePackage) error {
	temp, err := mod.ChoosePackage(data.IdPackage)
	if err != nil {
		return err
	}
	tempo, erro := mod.ListOrderData(data.IdOrder)
	if erro != nil {
		return erro
	}
	if len(tempo) != 1 {
		return errors.NotResolverErr
	}
	if tempo[0].Mass >= temp.LowerMass && tempo[0].Mass <= temp.UpperMass {
		dealErr := mod.ChangePackge(data)
		if dealErr != nil {
			fmt.Println("Данная упаковка не подходит")
			return dealErr
		}
		return nil
	}
	return errors.NotResolverErr
}

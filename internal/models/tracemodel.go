package models

import "time"

type DataUnitJson struct {
	IdOrder       uint32    `json:"idOrder"`
	IdUser        uint32    `json:"idUser"`
	IdPackage     uint32    `json:"idPackage"`
	DeliveredDate time.Time `json:"deliveredDate"`
	ReceivedDate  time.Time `json:"receivedDate"`
	DeadLine      time.Time `json:"deadLine"`
	RefundDate    time.Time `json:"refundDate"`
	Mass          uint32
}
type PositionRepoFilter struct {
	IdOrders []uint32 `json:"idOrders"`
}

func (f *PositionRepoFilter) Empty() bool {
	return f == nil || len(f.IdOrders) == 0
}

type ListContactFilter struct {
	PositionNames []string
}

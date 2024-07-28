package models

import "time"

type DataUnit struct {
	IdOrder       uint32 //ID заказа
	IdUser        uint32 //ID пользователя
	IdPackage     uint32
	DeliveredDate time.Time // Дата доставки на пвз
	ReceivedDate  time.Time // Дата получения
	DeadLine      time.Time // Срок день хранения
	RefundDate    time.Time // День возврата
	Mass          uint32
}
type PackageUnit struct {
	IdPackage   uint32
	PackageCost uint32
	PackageName string
	LowerMass   uint32
	UpperMass   uint32
}
type ChangePackage struct {
	IdOrder   uint32
	IdPackage uint32
}
type ReceiveOrderDeliver struct {
	IdOrder   uint32 //ID заказа
	IdUser    uint32 //ID пользователя
	IdPackage uint32
	DeadLine  int // Срок день хранения
	Mass      uint32
}
type RefundOrderDeliver struct {
	IdOrder uint32 //ID заказа
}
type ReceiveOrderUser struct {
	IdOrder uint32
}
type OrderList struct {
	IdUser uint32 //ID пользователя
}
type RefundUser struct {
	IdOrder uint32
	IdUser  uint32
}
type RefundList struct {
	Num uint32
}

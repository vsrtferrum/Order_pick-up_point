package storage

import (
	"time"
)

type OrderRecord struct {
	IdOrder       uint32    `db:"id_order"` //ID заказа
	IdUser        uint32    `db:"id_user"`  //ID пользователя
	IdPackage     uint32    `db:"id_package"`
	DeliveredDate time.Time `db:"delivered_date"` // Дата доставки на пвз
	ReceivedDate  time.Time `db:"recieved_date"`  // Дата получения
	DeadLine      time.Time `db:"dead_line"`      // Срок день хранения
	RefundDate    time.Time `db:"refund_date"`    // День возврата
	Mass          uint32    `db:"mass"`
}
type packageRecord struct {
	IdPackage   uint32 `db:"id_package"`
	PackageCost uint32 `db:"package_cost"`
	PackageName string `db:"package_name"`
	LowerMass   uint32 `db:"lower_mass"`
	UpperMass   uint32 `db:"upper_mass"`
}

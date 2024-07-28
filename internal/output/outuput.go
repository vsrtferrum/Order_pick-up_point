package output

import (
	"fmt"

	"gitlab.ozon.dev/berkinv/homework/internal/models"
)

func ListOrders(p []models.DataUnit, pack []models.PackageUnit, sz uint64) {
	hashT := make(map[uint32]models.PackageUnit)
	for _, u := range pack {
		hashT[u.IdPackage] = u
	}
	for i := uint64(0); i < uint64(len(p)); i += sz {
		fmt.Printf("IdOrder\tIdUser\tDeliveredTime\tRefundDeliverDate\t"+
			"AcceptByUserDate\tRefundByUserDate\tPackageName\tPackageCost [%d]-[%d]\n",
			i, min(uint64(len(p)), i+sz))
		for j := i; j < min(i+sz, uint64(len(p))); j++ {
			fmt.Println(p[j].IdOrder, "\t", p[j].IdUser, "\t",
				p[j].DeliveredDate.Format("2006-01-02 15:04:05"), "\t",
				p[j].DeadLine.Format("2006-01-02 15:04:05"), "\t",
				p[j].ReceivedDate.Format("2006-01-02 15:04:05"), "\t",
				p[j].RefundDate.Format("2006-01-02 15:04:05"), "\t",
				hashT[p[j].IdPackage].PackageName, "\t", hashT[p[j].IdUser].PackageCost)
			if j == uint64(len(p)-1) {
				fmt.Println("Выход из раздела")
				return
			}
		}
		fmt.Println("Введите q для выхода или любую другую комбинацию клавиш")
		var cmd string
		_, errScan := fmt.Scanln(&cmd)
		if cmd == "q" || errScan != nil {
			fmt.Println("Выход из раздела")
			return
		}
	}
	fmt.Println("Выход из раздела")
}

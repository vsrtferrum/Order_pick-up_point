package cli

import (
	"sync"

	log "gitlab.ozon.dev/berkinv/homework/internal/handlers/log"
)

func NewCLI(broker []string, topic string, mode bool) CLI {
	return CLI{
		workers: workers{
			wg:          &sync.WaitGroup{},
			numWorkers:  2,
			syncChan:    make(chan struct{}, 1),
			commandChan: make(chan worker, 1),
		},
		mode: mode,
		Log:  log.NewLogger(mode, broker, topic),
		commandList: []command{
			{
				name: help,
				description: "Команды :\n" +
					"help - Справка\n",
			},
			{
				name: acceptOrder,
				description: "accept - Принять заказ от курьера\t" +
					"Данные для ввода: id_заказа id_получателя время_хранения(дни)\n",
			},
			{
				name:        refundDeliver,
				description: "refund_d - Возврат заказа курьеру\tДанные для ввода: id_заказа\n",
			},
			{
				name: issueUser,
				description: "issue -  Выдать заказ клиенту\t" +
					"Данные для ввода: id_заказов через_пробел\n",
			},
			{
				name: listOrders,
				description: "list_o -  Получить список закзаов клиентов\t" +
					"Данные для ввода: id_клиента тип_выдачи " +
					"(число_последних_заказов (по умиолчанию - все готовые к выдаче))\n",
			},
			{
				name: refundUser,
				description: "refund_u - Получить возврат от клиента\t" +
					"Данные для ввода: id_заказа id_клиента\n",
			},
			{
				name: listRefund,
				description: "list_r - Получить список возвратов\t" +
					"Данные выдаются по 10 \n",
			},
			{
				name: setWorkersNum,
				description: "set-workers - Установить число горутин\t" +
					"Данные для ввода: числогорутин\n",
			},
			{
				name: addPackage,
				description: "add-package - Добавить новый вид упаковки\t" +
					"Данные для ввода: Цена_упаковки Название_упкаовки " +
					"нижняя_граница_массы вернхняя_граница_массы\n",
			},
			{
				name: changePackagetype,
				description: "change - Изменить вид упаковки\tДанные для ввода: id_заказа " +
					"id_нового_типа_упаковки\n",
			},
		},
	}
}

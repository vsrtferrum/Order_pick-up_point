package errors

import "errors"

var (
	NotFoundCommandErr           = errors.New("Введите команду заново")
	NotResolverErr               = errors.New("Фрагментарное несовпадение записей")
	OpenErr                      = errors.New("Источник данных поврежден")
	ExitErr                      = errors.New("Чтение прервано")
	GorutineCreateErr            = errors.New("Невозможно создать нужное колличество горутин")
	CantResolvArgsErr            = errors.New("Неверная внутренняя форма команды")
	CantRefundDeliverErr         = errors.New("Невозможно вернуть заказ курьеру")
	CantRecieveOrderUserErr      = errors.New("Невозможно выдать заказ пользователю")
	RefundUserErr                = errors.New("Невозможно принять возврат от клиента")
	CreationSyncKafkaProduserErr = errors.New("Невозможно создать kafka producer ")
	CloseSyncKafkaProduserErr    = errors.New("Невозможно закрыть kafka producer ")
	UnmarshalJsonErr             = errors.New("Невозможно демаршалировать json")
	FoundOrderEr                 = errors.New("Не удалось найти заказ")
	SubscribeErr                 = errors.New("Не удалось подписаться на партицию")
	InternalServerError          = errors.New("Ошибка сервера")
)

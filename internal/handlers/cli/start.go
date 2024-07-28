package cli

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
)

func (cli CLI) Start(commandName string, args []string) error {
	err := cli.Log.Input(commandName, args)
	if err != nil {
		return err
	}
	switch commandName {
	case help:
		return cli.help()
	case acceptOrder:
		return cli.acceptOrder(args)
	case refundDeliver:
		return cli.refundDeliver(args)
	case issueUser:
		return cli.issueUser(args)
	case listOrders:
		return cli.listOrders(args)
	case refundUser:
		return cli.refundUser(args)
	case listRefund:
		return cli.listRefund(args)
	case setWorkersNum:
		_, err := cli.setWorkersNum(args)
		return err
	case addPackage:
		return cli.addPackage(args)
	case changePackagetype:
		return cli.changePackage(args)
	case exit:
		return errors.ExitErr
	default:
		return errors.NotFoundCommandErr
	}
}

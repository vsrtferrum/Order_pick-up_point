package cli

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/input"
	"gitlab.ozon.dev/berkinv/homework/internal/output"
)

func (cli CLI) listOrders(arg []string) error {
	lo, size, err := input.CliListOrder(arg)
	if err != nil {
		return errors.CantResolvArgsErr
	}
	res, dealError := cli.Keepnthrow.OrderList(lo)
	if dealError != nil {
		return errors.OpenErr
	}
	if size == 0 {
		size = uint64(len(res))
	}
	pack, errpack := cli.Keepnthrow.ListPackage()
	if errpack != nil {
		return errors.OpenErr
	}
	output.ListOrders(res, pack, size)
	return nil
}

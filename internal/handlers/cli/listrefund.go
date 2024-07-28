package cli

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/input"
	"gitlab.ozon.dev/berkinv/homework/internal/output"
)

func (cli CLI) listRefund(arg []string) error {
	lr, err := input.CliListRefund(arg)
	if err != nil {
		return errors.CantResolvArgsErr
	}
	res, dealError := cli.Keepnthrow.RefundList()
	if dealError != nil {
		return errors.OpenErr
	}
	pack, errpack := cli.Keepnthrow.ListPackage()
	if errpack != nil {
		return errors.OpenErr
	}
	output.ListOrders(res, pack, lr)
	return nil
}

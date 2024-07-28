package cli

import (
	"fmt"

	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/input"
)

func (cli CLI) refundUser(arg []string) error {
	ru, err := input.CliRefundUser(arg)
	if err != nil {
		return errors.CantResolvArgsErr
	}
	dealError := cli.Keepnthrow.RefundUser(ru)
	fmt.Println(dealError)
	return nil
}

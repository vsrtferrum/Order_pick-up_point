package cli

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/input"
)

func (cli CLI) refundDeliver(arg []string) error {
	rd, err := input.CliRefundDeliver(arg)
	if err != nil {
		return errors.CantResolvArgsErr
	}
	dealError := cli.Keepnthrow.RefundDeliver(rd)
	return dealError
}

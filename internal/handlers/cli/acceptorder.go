package cli

import (
	"fmt"

	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/input"
	"gitlab.ozon.dev/berkinv/homework/internal/responses"
)

func (cli CLI) acceptOrder(arg []string) error {
	ao, err := input.CliAccept(arg)
	if err != nil {
		return errors.CantResolvArgsErr
	}
	writeError := cli.Keepnthrow.ReceiveOrderDeliver(ao)
	if writeError != nil {
		fmt.Println(responses.AcceptOk)
		return writeError
	}
	fmt.Println(responses.AcceptNo)
	return nil
}

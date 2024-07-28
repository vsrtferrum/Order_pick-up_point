package cli

import (
	"fmt"

	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/input"
	"gitlab.ozon.dev/berkinv/homework/internal/responses"
)

func (cli CLI) issueUser(arg []string) error {
	iu, err := input.CliIssueUser(arg)
	if err != nil {
		return errors.CantResolvArgsErr
	}
	dealError := cli.Keepnthrow.ReceiveOrderUser(iu)
	if dealError == nil {
		fmt.Println(responses.IssueOk)
	} else {
		fmt.Println(responses.IssueNo)
		return errors.CantResolvArgsErr
	}
	return nil
}

package cli

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/input"
)

func (cli CLI) addPackage(arg []string) error {
	ap, err := input.CliAddPackage(arg)
	if err != nil {
		return errors.CantResolvArgsErr
	}
	dealErr := cli.Keepnthrow.AddPackage(ap)
	if dealErr != nil {
		return errors.OpenErr
	}
	return nil
}

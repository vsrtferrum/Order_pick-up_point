package cli

import (
	"gitlab.ozon.dev/berkinv/homework/internal/errors"
	"gitlab.ozon.dev/berkinv/homework/internal/handlers/input"
)

func (cli CLI) changePackage(arg []string) error {
	cp, err := input.CliChangePack(arg)
	if err != nil {
		return errors.CantResolvArgsErr
	}
	errD := cli.Keepnthrow.ChangePackage(cp)
	if errD != nil {
		return errors.CantResolvArgsErr
	}
	return nil
}

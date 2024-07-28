package cli

import (
	"os"

	log "gitlab.ozon.dev/berkinv/homework/internal/handlers/log"
	"gitlab.ozon.dev/berkinv/homework/internal/module"
)

type Module interface {
	receiveOrderDeliver(arg []string) error
	refundDeliver(arg []string) error
	receiveOrderUser(arg []string) error
	orderList(arg []string) error
	refundUser(arg []string) error
	refundList(arg []string) error
	setWorkersNum(arg []string) error
	addPackage(arg []string) error
	changePackage(arg []string) error
	Start(commandName string, args []string) error
	Run() error
}

type CLI struct {
	Module
	mode        bool
	commandList []command
	Keepnthrow  module.Module
	workers     workers
	Log         *log.Logger
}

func (cli CLI) Run() error {
	args := os.Args[1:]
	if len(args) == 0 {
		cli.StartWorkers()
	} else {
		if len(args) < 2 {
			args = append(args, "")
		}
		err := cli.Start(args[0], args[1:])
		if err != nil {
			return err
		}
	}
	return nil
}

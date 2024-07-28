package cli

import "gitlab.ozon.dev/berkinv/homework/internal/handlers/input"

func (cli CLI) setWorkersNum(arg []string) (uint64, error) {
	sw, err := input.CliSetWorkersNum(arg)
	if err != nil {
		return 0, err
	}
	cli.stopAllWorkers(cli.workers.numWorkers)
	cli.workers.numWorkers = sw
	cli.startAllWorkers(cli.workers.numWorkers)
	return sw, nil
}

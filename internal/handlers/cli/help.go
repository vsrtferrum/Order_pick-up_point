package cli

import "fmt"

func (cli CLI) help() error {
	for _, iter := range cli.commandList {
		fmt.Print(iter.description)
	}
	return nil
}

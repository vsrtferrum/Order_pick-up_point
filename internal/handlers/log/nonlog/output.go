package nonlog

import (
	"fmt"
	"time"
)

func (nonlog *Nonlog) Input(command string, args []string) {
	argsString := fmt.Sprint(args)
	fmt.Printf("Command: %s\tCommndArgs:%s\tTime:%s\n", command, argsString, time.Now().Format("2006-01-02 15:04:05"))
}

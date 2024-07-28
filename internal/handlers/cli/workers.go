package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"gitlab.ozon.dev/berkinv/homework/internal/errors"
)

type worker struct {
	args []string
}
type workers struct {
	numWorkers  uint64
	commandChan chan worker
	wg          *sync.WaitGroup
	syncChan    chan struct{}
}

func (cli CLI) worker(indx int) {
	defer cli.workers.wg.Done()
	for {
		select {
		case <-cli.workers.syncChan:
			fmt.Printf("\nРаботник №%d заврешился \n", indx)
			return
		case work, ok := <-cli.workers.commandChan:
			if !ok {
				continue
			}
			fmt.Printf("\nРаботник №%d приступил к работе \n", indx)
			if len(work.args) == 1 {
				work.args = append(work.args, "")
			}
			errStart := cli.Start(work.args[0], work.args[1:])
			if errStart == nil {
				fmt.Printf("\nРаботник № %d работу окончил\n", indx)
			}
		}
	}
}

func (cli CLI) StartWorkers() {
	defer close(cli.workers.syncChan)
	scanner := bufio.NewScanner(os.Stdin)
	cli.startAllWorkers(cli.workers.numWorkers)
	contxt, backup := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	exec := make(chan os.Signal, 1)
	signal.Notify(exec, syscall.SIGINT, syscall.SIGTERM)
	go func(contxt context.Context) {
		for {
			select {
			case <-contxt.Done():
				fmt.Println("Завершение работы")
				close(cli.workers.commandChan)
				return
			default:
				scanner.Scan()
				text := scanner.Text()
				args := strings.Fields(text)
				if len(args) == 0 {
					continue
				}
				if args[0] == setWorkersNum {
					num, err := cli.setWorkersNum(args[1:])
					if err != nil {
						fmt.Println(errors.GorutineCreateErr)
					}
					cli.workers.numWorkers = num
					continue
				}
				if args[0] == exit {
					exec <- syscall.SIGINT
				}
				worker := worker{args: args}
				cli.workers.commandChan <- worker
			}
		}
	}(contxt)
	<-exec
	backup()
	fmt.Println("Завершение горутин")
	cli.stopAllWorkers(cli.workers.numWorkers)
	cli.workers.wg.Wait()
}
func (cli *CLI) startAllWorkers(count uint64) {
	fmt.Println("Инициализация горутин")
	for i := 1; i < int(count); i++ {
		cli.workers.wg.Add(1)
		go cli.worker(i)
	}
	fmt.Println("Инициализация горутин завршена")
}
func (cli *CLI) stopAllWorkers(count uint64) {
	fmt.Println("Завершение горутин")
	for i := 1; i < int(count); i++ {
		cli.workers.syncChan <- struct{}{}
	}
	fmt.Println("Горутины завершены")
}

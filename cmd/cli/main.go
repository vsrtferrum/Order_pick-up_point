package main

import (
	"flag"

	"gitlab.ozon.dev/berkinv/homework/internal/handlers"
	cli "gitlab.ozon.dev/berkinv/homework/internal/handlers/cli"
)

var (
	fulldbreq string
	mode      string
	brokers   = []string{
		"127.0.0.1:9091",
	}
	topic string = "Log"
)

func main() {
	flag.StringVar(&fulldbreq, "fulldbreq", "user=vsrtf dbname=postgres  sslmode=disable", "fulldbreq")
	flag.StringVar(&mode, "mode", "kafka", "mode")
	client := cli.NewCLI(brokers, topic, handlers.IsKafkaMode(mode))

	err := client.Keepnthrow.SetFullDatabaseReq(fulldbreq)
	if err != nil {
		panic(err)
	}

	err = client.Run()
	if err != nil {
		panic(err)
	}
}

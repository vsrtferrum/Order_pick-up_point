package handlers

const (
	kafka   = "kafka"
	console = "console"
)

func IsKafkaMode(mod string) bool {
	switch mod {
	case kafka:
		return true
	case console:
		return false
	default:
		return true
	}

}

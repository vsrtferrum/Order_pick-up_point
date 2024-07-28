package log

func (log *Log) Input(command string, args []string) error {
	msg := log.sender.CreateMessage(command, args)
	err := log.sender.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

package models

import "time"

type LogMessage struct {
	AnswerID    int
	Status      string
	CreatedAt   time.Time
	CommandName string
	Args        []string
}

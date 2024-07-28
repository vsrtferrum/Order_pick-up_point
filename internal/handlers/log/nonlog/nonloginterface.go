package nonlog

import "gitlab.ozon.dev/berkinv/homework/internal/models"

type NonLogInterface interface {
	Input(input models.LogMessage)
}
type Nonlog struct {
	NonLogInterface
}

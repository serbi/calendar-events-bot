package calendar

import (
	"time"
)

type Request struct {
	Id      string
	TimeMin time.Time
	TimeMax time.Time
}

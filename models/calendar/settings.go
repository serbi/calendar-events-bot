package calendar

import "os"

var (
	Token      = os.Getenv("CALENDAR_TOKEN")
	TimeZone   = "Europe/Warsaw"
	SalaSofaId = os.Getenv("CALENDAR_SOFA_ID")
)

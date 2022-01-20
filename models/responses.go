package models

import "time"

type Response struct {
	Sign ZodiacSign
	Date time.Time
	Text string
}

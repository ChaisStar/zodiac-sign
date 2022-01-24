package models

import "time"

type Response struct {
	Sign  ZodiacSign
	Date  time.Time
	Texts map[string]string
}

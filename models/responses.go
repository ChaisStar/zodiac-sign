package models

import "time"

type Response struct {
	Sign  int64
	Date  time.Time
	Texts map[string]string
	Type  HoroscopeType
}

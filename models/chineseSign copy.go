package models

//go:generate stringer -type=HoroscopeType
type HoroscopeType int64

const (
	Zodiac HoroscopeType = iota
	Chinese
)

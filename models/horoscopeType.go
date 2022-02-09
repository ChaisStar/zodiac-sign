package models

//go:generate stringer -type=HoroscopeType
type HoroscopeType int64

const (
	FrenchDefault HoroscopeType = iota
	FrenchChinese
	YahooCommon
)

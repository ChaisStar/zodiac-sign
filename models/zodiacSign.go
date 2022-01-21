package models

//go:generate stringer -type=ZodiacSign
type ZodiacSign int64

const (
	Aries ZodiacSign = iota + 1
	Taurus
	Gemini
	Cancer
	Leo
	Virgo
	Libra
	Scorpio
	Sagittarius
	Capricorn
	Aquarius
	Pisces
)

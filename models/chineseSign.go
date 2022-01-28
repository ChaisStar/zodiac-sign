package models

//go:generate stringer -type=ChineseSign
type ChineseSign int64

const (
	Rat ChineseSign = iota + 1
	Ox
	Tiger
	Rabbit
	Dragon
	Snake
	Horse
	Goat
	Monkey
	Rooster
	Dog
	Pig
)

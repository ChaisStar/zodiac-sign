package models

type Request struct {
	Sign      []ZodiacSign
	StartDate string
	EndDate   string
}

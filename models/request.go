package models

type Request struct {
	Signs     []ZodiacSign
	StartDate string
	EndDate   string
}

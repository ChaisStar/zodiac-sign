package models

type TypesArray []HoroscopeType

func (types TypesArray) Has(horoscopeType HoroscopeType) bool {
	for _, b := range types {
		if b == horoscopeType {
			return true
		}
	}
	return false
}

type Request struct {
	StartDate string
	EndDate   string
	Types     TypesArray
	Signs     []int64
}

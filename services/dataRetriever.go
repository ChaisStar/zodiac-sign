package services

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ChaisStar/zodiac-sign/models"

	"github.com/PuerkitoBio/goquery"
)

func CreateFormData(zodiacSign models.ZodiacSign, date time.Time) url.Values {
	return url.Values{
		"sign": {fmt.Sprint(int(zodiacSign))},
		"jour": {fmt.Sprint(date.Day())},
		"mois": {fmt.Sprint(int(date.Month()))},
		"an":   {fmt.Sprint(date.Year())},
	}
}

func GetHtml(formData url.Values) *goquery.Document {
	resp, err := http.PostForm("https://www.asiaflash.com/horoscope/parjour_365_jours.php", formData)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return doc
}

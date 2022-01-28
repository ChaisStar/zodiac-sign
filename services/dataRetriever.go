package services

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func CreateFormData(sign int64, date time.Time) url.Values {
	return url.Values{
		"sign": {fmt.Sprint(sign)},
		"jour": {fmt.Sprint(date.Day())},
		"mois": {fmt.Sprint(int(date.Month()))},
		"an":   {fmt.Sprint(date.Year())},
	}
}

func GetZodiacHtml(formData url.Values) *goquery.Document {
	return getHtml("https://www.asiaflash.com/horoscope/parjour_365_jours.php", formData)
}

func GetChineseHtml(formData url.Values) *goquery.Document {
	return getHtml("https://www.asiaflash.com/horoscope/cjour_365_jours.php", formData)
}

func getHtml(url string, formData url.Values) *goquery.Document {
	resp, err := http.PostForm(url, formData)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return doc
}

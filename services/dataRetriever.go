package services

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ChaisStar/zodiac-sign/models"
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

func GetFrenchDefaultHtml(formData url.Values) *goquery.Document {
	return getFranchHtml("https://www.asiaflash.com/horoscope/parjour_365_jours.php", formData)
}

func GetFrenchChineseHtml(formData url.Values) *goquery.Document {
	return getFranchHtml("https://www.asiaflash.com/horoscope/cjour_365_jours.php", formData)
}

func GetYahooCommonHtml(sign models.ZodiacSign, date time.Time) *goquery.Document {
	url := "https://www.yahoo.com/lifestyle/horoscope/" + strings.ToLower(sign.String()) + "/daily-" + date.Format("20060102") + ".html"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return doc
}

func getFranchHtml(url string, formData url.Values) *goquery.Document {
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

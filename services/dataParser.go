package services

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var supportedZodiacTypes = []string{
	"Santé",
	"Travail",
	"Famille",
}

var supportedChineseTypes = []string{
	"Vie sociale",
}

func ParseChineseHtml(html *goquery.Document) map[string]string {
	return parseHtml(html, supportedChineseTypes)
}

func ParseZodiacHtml(html *goquery.Document) map[string]string {
	return parseHtml(html, supportedZodiacTypes)
}

func parseHtml(html *goquery.Document, supportedTypes []string) map[string]string {
	// var sb strings.Builder
	result := make(map[string]string)

	html.Find(".af_rubrique p").Not("form").Each(func(i int, s *goquery.Selection) {
		var span string

		s.Find("span").Each(func(i2 int, s2 *goquery.Selection) {
			span = s2.Text()
		})

		supported := ""

		for _, supportedType := range supportedTypes {
			if strings.Contains(span, supportedType) {
				supported = supportedType
				break
			}
		}

		if supported == "" {
			return
		}

		d := strings.Replace(s.Text(), span, "", -1)

		if !strings.Contains(span, "●") {
			return
		}

		result[supported] = d
	})

	return result
}

package services

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var supportedFrenchZodiacTypes = []string{
	"Santé",
	"Travail",
	"Famille",
}

var supportedFrenchChineseTypes = []string{
	"Vie sociale",
}

var supportedYahooCommonTypes = []string{
	"daily",
}

func ParseFrenchChineseHtml(html *goquery.Document) map[string]string {
	return parseFrenchHtml(html, supportedFrenchChineseTypes)
}

func ParseFrenchZodiacHtml(html *goquery.Document) map[string]string {
	return parseFrenchHtml(html, supportedFrenchZodiacTypes)
}

func ParseYahooCommonHtml(html *goquery.Document) map[string]string {
	result := make(map[string]string)
	html.Find(".Horoscope div div div ul a").Each(func(i int, s *goquery.Selection) {
		span := s.Text()
		supported := ""
		for _, supportedType := range supportedYahooCommonTypes {
			if strings.Contains(span, supportedType) {
				supported = supportedType
				break
			}
		}
		if supported == "" {
			return
		}

		k := s.Parent().Parent().Parent().Parent().Children().Last().Children().First().Text()
		result["Common"] = strings.Split(k, "<a")[0]
	})

	return result
}

func ParseAstrolisDetailedHtml(html *goquery.Document) map[string]string {
	result := make(map[string]string)
	html.Find("span[itemprop='articleBody']").Each(func(i int, s *goquery.Selection) {
		result["Detailed"] = strings.TrimSpace(s.Text())
	})

	return result
}

func ParseAstrolisLoveHtml(html *goquery.Document) map[string]string {
	result := make(map[string]string)
	html.Find("span[itemprop='articleBody']").Each(func(i int, s *goquery.Selection) {
		result["Love"] = strings.Split(strings.TrimSpace(s.Text()), "<br>")[0]
	})

	return result
}

func parseFrenchHtml(html *goquery.Document, supportedTypes []string) map[string]string {
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

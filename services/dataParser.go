package services

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHtml(html *goquery.Document) string {
	var sb strings.Builder

	html.Find(".af_rubrique p").Not("form").Each(func(i int, s *goquery.Selection) {

		var span string

		s.Find("span").Each(func(i2 int, s2 *goquery.Selection) {
			span = s2.Text()
			sb.WriteString(fmt.Sprintf("%v ", strings.Trim(span, " ●")))
		})

		d := strings.Replace(s.Text(), span, "", -1)

		if !strings.Contains(span, "●") {
			return
		}

		sb.WriteString(d)
	})

	return sb.String()
}

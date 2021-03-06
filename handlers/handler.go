package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ChaisStar/zodiac-sign/models"
	"github.com/ChaisStar/zodiac-sign/services"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var request models.Request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	response := getResponse(request)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.xlsx\"", request.StartDate))
	w.Write(response)
}

func getResponse(request models.Request) []byte {
	startDate := parseDate(request.StartDate)
	endDate := parseDate(request.EndDate)

	builder := services.NewBuilder()

	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		for _, sign := range request.Signs {
			formData := services.CreateFormData(sign, date)

			if request.Types.Has(models.FrenchChinese) {
				html := services.GetFrenchChineseHtml(formData)
				data := services.ParseFrenchChineseHtml(html)
				response := models.Response{Sign: sign, Date: date, Texts: data, Type: models.FrenchChinese}
				builder.Add(response)
			}

			if request.Types.Has(models.FrenchDefault) {
				html := services.GetFrenchDefaultHtml(formData)
				data := services.ParseFrenchZodiacHtml(html)
				response := models.Response{Sign: sign, Date: date, Texts: data, Type: models.FrenchDefault}
				builder.Add(response)
			}

			if request.Types.Has(models.YahooCommon) {
				html := services.GetYahooCommonHtml(models.ZodiacSign(sign), date)
				data := services.ParseYahooCommonHtml(html)
				response := models.Response{Sign: sign, Date: date, Texts: data, Type: models.YahooCommon}
				builder.Add(response)
			}

			if request.Types.Has(models.AstrolisDetailed) {
				html := services.GetAstrolisDetailedHtml(models.ZodiacSign(sign), date)
				data := services.ParseAstrolisDetailedHtml(html)
				response := models.Response{Sign: sign, Date: date, Texts: data, Type: models.AstrolisDetailed}
				builder.Add(response)
			}

			if request.Types.Has(models.AstrolisLove) {
				html := services.GetAstrolisLoveHtml(models.ZodiacSign(sign), date)
				data := services.ParseAstrolisLoveHtml(html)
				response := models.Response{Sign: sign, Date: date, Texts: data, Type: models.AstrolisLove}
				builder.Add(response)
			}
		}
	}

	return builder.Bytes()
}

func parseDate(dateString string) time.Time {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		log.Fatalln(err)
	}

	return date
}

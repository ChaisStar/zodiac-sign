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
			html := services.GetHtml(formData)
			data := services.ParseHtml(html)
			response := models.Response{Sign: sign, Date: date, Text: data}
			builder.Add(response)
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

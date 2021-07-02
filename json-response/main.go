package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type body struct {
	PublishingOffice string `json:"publishingOffice"`
	ReportDatetime   string `json:"reportDatetime"`
	TargetArea       string `json:"targetArea"`
	HeadlineText     string `json:"headlineText"`
	Text             string `json:"text"`
}

func main() {
	res, err := http.Get(`https://www.jma.go.jp/bosai/forecast/data/overview_forecast/130000.json`)
	if err != nil {
		log.Fatalf(`unexpeted error: %s`, err)
	}
	var reader io.Reader
	reader = res.Body
	// comment out out
	// reader = io.TeeReader(reader, os.Stdout)
	var v body
	if err := json.NewDecoder(reader).Decode(&v); err != nil {
		log.Fatalf(`unexpeted error: %s`, err)
	}
}

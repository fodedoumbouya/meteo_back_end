package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"meteo_back_end/models"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func GetDataURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// data :=models.ResponseRquest{

	// }
	fmt.Println("innnnnnnnnn")
	msg := "error"
	code := 400
	var data models.Temperature
	//r.URL.Host
	if len(id) > 0 {
		data = GetDataFromUrl(id)
		code = 200
		msg = "sucess"
	}

	resp := models.ResponseRquest{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	json.NewEncoder(w).Encode(resp)

}

func GetDataFromUrl(id string) models.Temperature {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	url := fmt.Sprintf("http://localhost:8085/widget?id=%v", id)
	// navigate to URL
	var html string
	var example string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("#allmeteo-temperature", chromedp.ByID),
		chromedp.EvaluateAsDevTools(`document.documentElement.outerHTML`, &html),
		//chromedp.InnerHTML("#allmeteo-temperature", &example, chromedp.ByID),

	); err != nil {
		panic(err)
	}
	fmt.Println(example)

	//fmt.Println(html)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	temperature := doc.Find("#allmeteo-temperature").Text()
	humidity := doc.Find("#allmeteo-humidity").Text()

	fmt.Printf("Temperature: %s\n", temperature)
	fmt.Printf("Humidity: %s\n", humidity)

	return models.Temperature{
		Temperature: temperature,
		Humidity:    humidity,
	}

}

package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"meteo_back_end/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func GetDataURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// data :=models.ResponseRquest{

	// }
	msg := "error"
	code := 400
	var data string
	//models.Temperature
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

func GetDataFromUrl(id string) string {
	// create context

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	url := fmt.Sprintf("http://localhost:8085/api/widget?id=%v", id)
	// navigate to URL
	var html string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("#allmeteo-temperature", chromedp.ByID),
		chromedp.EvaluateAsDevTools(`document.documentElement.outerHTML`, &html),
		//chromedp.InnerHTML("#allmeteo-temperature", &example, chromedp.ByID),

	); err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	temperature, _ := strconv.ParseFloat(doc.Find("#allmeteo-temperature").Text(), 64)
	//allmeteo-deworfrost
	humidity, _ := strconv.ParseFloat(doc.Find("#allmeteo-humidity").Text(), 64)
	// deworfrost, _ := strconv.ParseFloat(doc.Find("#allmeteo-deworfrost").Text(), 64)
	dewPoint, _ := strconv.ParseFloat(doc.Find("#allmeteo-dewPoint").Text(), 64)
	rain, _ := strconv.ParseFloat(doc.Find("#allmeteo-rain").Text(), 64)
	// pressure, _ := strconv.ParseFloat(doc.Find("#allmeteo-pressure").Text(), 64)
	irradiation, _ := strconv.ParseFloat(doc.Find("#allmeteo-irradiation").Text(), 64)
	wetbulb, _ := strconv.ParseFloat(doc.Find("#allmeteo-wetbulb").Text(), 64)

	fmt.Printf("Temperature: %s\n", temperature)
	fmt.Printf("Humidity: %s\n", humidity)

	return getWeatherType(humidity, temperature, rain, irradiation, dewPoint, wetbulb)

}

func getWeatherType(humidity, temperature, rain, irradiation, dewPoint, wetBulb float64) string {
	// humidity := 39.4   // en %
	// temperature := 18  // en °C
	// rain := 0          // en mm
	// irradiation := 718 // en W/m2
	// dewPoint := 3.6    // en °C
	// wetBulb := 10.2    // en °C

	var typ string

	if temperature >= 30 && irradiation >= 500 {
		typ = "scorchingSun"
	} else if temperature >= 15 && temperature < 30 && humidity <= 70 && rain < 5 {
		typ = "sunset"
	} else if temperature <= 0 && (dewPoint < 0 || wetBulb < 0) && rain > 0 {
		typ = "frosty"
	} else if temperature <= 0 && rain > 0 {
		typ = "snowfall"
	} else if temperature <= 5 && (rain > 0 || (snowfall(dewPoint, temperature) > 0)) {
		typ = "showerSleet"
	} else if irradiation < 300 && rain > 0 {
		typ = "rainyOvercast"
	} else {
		typ = "stormy"
	}
	return typ
}

func snowfall(dewPoint float64, temperature float64) int {
	// Calcul de la différence entre le point de rosée et la température
	difference := temperature - dewPoint
	if difference < 2 {
		return 0
	} else if difference < 5 {
		return 1
	} else if difference < 10 {
		return 2
	} else {
		return 3
	}
}

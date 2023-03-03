package main

import (
	"encoding/csv"
	"fmt"
	"meteo_back_end/api"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"log"
	"os"

	"github.com/gocolly/colly"
)

func Handler() {
	http.HandleFunc("/station", api.GetStation)
	http.HandleFunc("/widget", api.GetWidget)

	fmt.Println("server running 8085")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		fmt.Println("server Failed")
	}
}

func scrapper() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{
				el.ChildText("td:nth-child(1)"),
				el.ChildText("td:nth-child(2)"),
				el.ChildText("td:nth-child(3)"),
			})
		})
		fmt.Println("Scrapping Complete")
	})
	c.Visit("https://www.w3schools.com/html/html_tables.asp")
}

func main() {
	Handler()

}

//
//http.HandleFunc("/data", getAllData)
//http.HandleFunc("/v", testHtml)
//type MeteoData struct {
//	Time             time.Time `json:"time"`
//	Battery_V        float64   `json:"battery_V"`
//	Wdir_Avg10       int64     `json:"wdir_Avg10"`
//	Wdir_Gust10      int64     `json:"wdir_Gust10"`
//	Wdir_Max10       int64     `json:"wdir_Max10"`
//	Wdir_Min10       int64     `json:"wdir_Min10"`
//	Wdir_Avg10_mPs   float64   `json:"wdir_Avg10_mPs"`
//	Wdir_Max10_mPs   float64   `json:"wdir_Max10_mPs"`
//	Wind_Stdev10_mPs float64   `json:"wind_Stdev10_mPs"`
//	Wdir_Stdev10     float64   `json:"wdir_Stdev10"`
//}
//
//type AllMeteoData struct {
//	Data []MeteoData
//}
//
//func (d *AllMeteoData) contient(t time.Time) (bool, MeteoData) {
//	var mData MeteoData
//	found := false
//	for _, v := range *&d.Data {
//		if v.Time == t {
//			mData = v
//			found = true
//			break
//		}
//	}
//	return found, mData
//}
//
//func getAllData(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	data := loadData()
//	// fmt.Println("lenght: ", len(data.Data))
//
//	startDate := time.Date(2022, 4, 12, 0, 0, 0, 0, time.UTC) // Start date
//	endDate := time.Date(2022, 4, 13, 0, 0, 0, 0, time.UTC)   // End date
//	dataFormTime := GetDataFromTime(data, startDate, endDate)
//	// fmt.Println("dataFormTime lenght: ", len(dataFormTime))
//	json.NewEncoder(w).Encode(dataFormTime)
//}
//func testHtml(w http.ResponseWriter, r *http.Request) {
//	//w.Header().Set("Content-Type", "application/html")
//	fmt.Fprintf(w,
//		`
//	<!DOCTYPE html>
//<html>
//<head>
//<title>HTML, CSS and JavaScript demo</title>
//</head>
//<body>
//<!-- Start your code here -->
//<div style="width: 400px; height:300px; border: 5px solid green;">
//  <div id="allm13" class="allmeteo-widget" data-ws="2108SW031">
//</div>
//</div>
//
//<!-- End your code here -->
//</body>
//  <script type="text/javascript" src="https://weather.allmeteo.com/widget/allmeteo.widget.js">  </script>
//</html>
//
//`,
//	)
//}
//
//func GetDataFromTime(data AllMeteoData, startDate time.Time, endDate time.Time) []MeteoData {
//	var dataFormTime []MeteoData
//
//	// Calculate the difference between start and end dates in minutes
//	diff := endDate.Sub(startDate).Minutes()
//	// Iterate over the difference in minutes and print the dates in 10 minute intervals
//	for i := 0; i <= int(diff); i += 10 {
//
//		date := startDate.Add(time.Duration(i) * time.Minute)
//		// fmt.Println(date)
//		resp, d := data.contient(date)
//		if resp {
//			dataFormTime = append(dataFormTime, d)
//		}
//
//	}
//
//	return dataFormTime
//
//}
//
//func loadData() AllMeteoData {
//	records, err := readData("data/meteoWindCorte.csv")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var dates []MeteoData
//
//	for _, record := range records {
//		_time, _ := time.Parse("2006-01-02 15:04:05", record[0])
//		_battery_V, _ := strconv.ParseFloat(record[1], 64)
//		_wdir_Avg10, _ := strconv.ParseInt(record[2], 36, 64)
//		_wdir_Gust10, _ := strconv.ParseInt(record[3], 36, 64)
//		_wdir_Max10, _ := strconv.ParseInt(record[4], 36, 64)
//		_wdir_Min10, _ := strconv.ParseInt(record[5], 36, 64)
//		_wdir_Avg10_mPs, _ := strconv.ParseFloat(record[6], 64)
//		_wdir_Max10_mPs, _ := strconv.ParseFloat(record[7], 64)
//		_wind_Stdev10_mPs, _ := strconv.ParseFloat(record[8], 64)
//		_wdir_Stdev10, _ := strconv.ParseFloat(record[9], 64)
//		data := MeteoData{
//			Time:             _time,
//			Battery_V:        _battery_V,
//			Wdir_Avg10:       _wdir_Avg10,
//			Wdir_Gust10:      _wdir_Gust10,
//			Wdir_Max10:       _wdir_Max10,
//			Wdir_Min10:       _wdir_Min10,
//			Wdir_Avg10_mPs:   _wdir_Avg10_mPs,
//			Wdir_Max10_mPs:   _wdir_Max10_mPs,
//			Wind_Stdev10_mPs: _wind_Stdev10_mPs,
//			Wdir_Stdev10:     _wdir_Stdev10,
//		}
//		dates = append(dates, data)
//
//		// fmt.Printf("%s %v is a %v\n", data.Time, data.Battery_V,
//		// 	data.wdir_Avg10)
//	}
//	return AllMeteoData{
//		Data: dates,
//	}
//}
//
//func readData(fileName string) ([][]string, error) {
//
//	f, err := os.Open(fileName)
//
//	if err != nil {
//		return [][]string{}, err
//	}
//
//	defer f.Close()
//
//	r := csv.NewReader(f)
//
//	// skip first line
//	if _, err := r.Read(); err != nil {
//		return [][]string{}, err
//	}
//
//	records, err := r.ReadAll()
//
//	if err != nil {
//		return [][]string{}, err
//	}
//
//	return records, nil
//}

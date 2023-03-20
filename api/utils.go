package api

import (
	"database/sql"
	"fmt"
	"meteo_back_end/constant"
	"meteo_back_end/models"
)

var (
	id                                          int
	serialNumber, deviceNumber, location, model string
	lat, log                                    float64
)

// getAllStationData return data and boolean to check if there is no error
func getAllStationData() ([]models.Station, bool) {
	stations := []models.Station{}

	db, _ := sql.Open("mysql", constant.MySqlConfig())

	defer db.Close()

	sqlQuery := fmt.Sprintf("SELECT * FROM %v.stations", constant.DBName)
	result, err := db.Query(sqlQuery)
	if CheckError(err) {
		fmt.Println("Check Error Failed: ", err)
		return stations, true
	}
	defer result.Close()

	for result.Next() {
		result.Scan(&id, &serialNumber, &deviceNumber, &location, &model, &lat, &log)
		stations = append(stations, models.Station{
			Id:           id,
			SerialNumber: serialNumber,
			DeviceNumber: deviceNumber,
			Location:     location,
			Model:        model,
			Lat:          lat,
			Log:          log,
		})
	}

	return stations, false

}

func getHtml(id string) string {
	html := fmt.Sprintf(
		`
	<html>
    <head>
        <style>
            html, body{
                height: 100%v;
                 font-size: 50px;
            }
            .parent > * {
                margin: 0 auto;
                 font-size: 50px;

            }
            .parent {
                width: 100%v; 
                height: 100%v;
                font-size: 60px;
            }
            .child {
                width: 800px; 
                height:1000px; 
                 font-size: 60px;
                color: red;

                
            }
        </style>
    </head>
    <body>
        <div class="parent">
            <div class="child">
                <div  class="allmeteo-widget" data-ws="%v"></div>
            </div>
        </div>
    </body>
    <script type="text/javascript" src="https://weather.allmeteo.com/widget/allmeteo.widget.js">  </script>
</html>

	`, "%", "%", "%", id,
	)

	return html
}

func CheckError(err error) bool {
	return err != nil
}

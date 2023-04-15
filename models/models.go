package models

type Station struct {
	Id           int     `josn:"id"`
	SerialNumber string  `json:"serialNumber"`
	DeviceNumber string  `json:"deviceNumber"`
	Location     string  `json:"location"`
	Model        string  `json:"model"`
	Lat          float64 `json:"lat"`
	Log          float64 `json:"log"`
}
type Temperature struct {
	Temperature string `josn:"temperature"`
	Humidity    string `josn:"humidity"`
}

type Connection struct {
	Host     string
	Port     string
	Password string
	DbName   string
}
type ResponseRquest struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
}

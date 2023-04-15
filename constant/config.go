package constant

import (
	"meteo_back_end/models"
)

var Mysql = models.Connection{
	Host:     "localhost:3306", //              //"8.134.122.202" //"8.208.76.170" // "localhost", //
	Password: "",               // "pass123",        //"1302313",  //"root",      //, "root",      //
	DbName:   DBName,
}

func MySqlConfig() string {
	//connectionParameters := ""
	//if DebugMode {
	//	connectionParameters = "root:" + Mysql.Password + "@tcp(" + Mysql.Host + ")/" + Mysql.DbName + ""
	//} else {
	//	connectionParameters = "root:" + mysqlRelase.Password + "@tcp(" + mysqlRelase.Host + ")/" + mysqlRelase.DbName + ""
	//}
	connectionParameters := "root:" + Mysql.Password + "@tcp(" + Mysql.Host + ")/" + Mysql.DbName + ""
	return connectionParameters
}

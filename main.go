package main

import (
	DBconfig "todo/connection"
	routes "todo/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = DBconfig.ConnectionDB()
)

func main() {
	//Đợi cho đến khi mở xong rồi đóng
	DBconfig.MigrateDB(db)
	defer DBconfig.DisconnectDB(db)

	routes.Routes()

}

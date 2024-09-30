package connection

//This class help connect and disconnect to database
import (
	"fmt"
	"todo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Details for connecting database
var (
	host     string = "localhost"
	port     string = "3306"
	username string = "root"
	password string = ""
	database string = "abc"
)

func ConnectionDB() *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil
	}

	// Kiểm tra kết nối OK hay chưa
	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil
	}

	return db
}
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Can't connect to database ,disconnect ")
	}
	dbSQL.Close()
}
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.Todo{})
	fmt.Println("Database Migration OK")
}

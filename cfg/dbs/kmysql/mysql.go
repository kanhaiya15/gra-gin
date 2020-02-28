package kmysql

import (
	"fmt"

	"github.com/kanhaiya15/gra-gin/models"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DB gorm.DB
var DB *gorm.DB

// DBConfig Db Config struct
type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// NewConfig generates a mysql configuration object which
// This mysql instance will becomes the single source of
// truth for the app configuration.
func NewConfig() {
	conn, err := getConnection()
	if err != nil {
		panic(err.Error())
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conn.Username, conn.Password, conn.Host, conn.Port, conn.DBName)

	DB, err = gorm.Open("mysql", connection)
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&models.TodoModel{})
}

func getConnection() (conn DBConfig, err error) {
	host := viper.GetString("APP.DB.MYSQL.HOST")
	port := viper.GetString("APP.DB.MYSQL.PORT")
	username := viper.GetString("APP.DB.MYSQL.USERNAME")
	password := viper.GetString("APP.DB.MYSQL.PASSWORD")
	dbName := viper.GetString("APP.DB.MYSQL.NAME")
	conn = DBConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbName,
	}
	return conn, nil
}

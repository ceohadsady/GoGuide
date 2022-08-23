package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var connectDB *gorm.DB
var err error

func ConnectDB() (*gorm.DB, error) {
	myDSN := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Vientiane",
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		viper.GetInt("db.password"),
		viper.GetString("db.database"),
		viper.GetString("db.port"),
	)
	fmt.Println("CONNECTING_TO_DB")
	connectDB, err = gorm.Open(postgres.Open(myDSN), &gorm.Config{
		NowFunc: func() time.Time {
			tim, _ := time.LoadLocation("Asia/Vientiane")
			return time.Now().In(tim)
		},
	})
	if err != nil {
		log.Fatal("CONNECT_DATABASE_ERROR", err)
	}
	fmt.Println("DB_CONNECTED")
	return connectDB, nil
}

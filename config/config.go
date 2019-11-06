package config

/*
	created by Ali Sadikin
	this file is configuration by viper that invoked config.json (env) in project directory
*/
import (
	"fmt"
	"os"
	"phonebook/general"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

/*
	-- checking mode view (development/production)
*/
func CheckConfiguration() {
	viperConfig()
	mode := viper.GetString("mode.name")
	if mode != "production" {
		fmt.Println("this service is under development mode")
		general.WriteErrorLog("this service is under development mode")
		os.Exit(3)
	}

}

/*
	-- database configuration (configuration could be seen in confing.json file)
*/
func ConnectDB() (*gorm.DB, error) {
	viperConfig()
	DBDriver := viper.GetString("database.driver")
	DBName := viper.GetString("database.name")
	DBUser := viper.GetString("database.user")
	DBPassword := viper.GetString("database.password")
	DBHost := viper.GetString("tcp.host")
	DBPort := viper.GetString("tcp.port")
	db, err := gorm.Open(DBDriver, DBUser+":"+DBPassword+"@tcp("+DBHost+":"+DBPort+")/"+DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		general.WriteErrorLog(err.Error())
	}

	db.LogMode(true)
	db.Debug()
	return db, nil
}

func viperConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		general.WriteErrorLog("the configuration may be not set clearly")
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

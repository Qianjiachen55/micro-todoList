package conf

import (
	"fmt"
	"github.com/Qianjiachen55/micro-todoList/model"
	"github.com/spf13/viper"
	"strings"
)

func Init()  {
	home := "conf/config"
	viper.AddConfigPath(home)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig();err != nil{
		fmt.Println("error:",err)
		panic(err)
	}

	//viper.Get("mysql.DB")
	path := LoadMysqlData()
	model.Database(path)

}

func LoadMysqlData() string {
	//db := viper.GetString("mysql.DB")
	passwd := viper.GetString("mysql.Passwd")
	host := viper.GetString("mysql.Host")
	port := viper.GetString("mysql.Port")
	user := viper.GetString("mysql.User")
	dbName := viper.GetString("mysql.DBName")

	return strings.Join([]string{user,":",passwd,"@tcp(",host,":",port,")/",dbName,"?charset=utf8&parseTime=true"},"")
}



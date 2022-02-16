package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"task/model"
	"strings"
)

func Init()  {
	home := "conf/"
	viper.AddConfigPath(home)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig();err != nil{
		fmt.Println("error:",err)
		panic(err)
	}

	//viper.Get("mysql.DB")
	LoadMysqlData()
	//model.Database(path)
	LoadRabbitMQ()

}


func LoadMysqlData()  {
	//db := viper.GetString("mysql.DB")
	passwd := viper.GetString("mysql.Passwd")
	host := viper.GetString("mysql.Host")
	port := viper.GetString("mysql.Port")
	user := viper.GetString("mysql.User")
	dbName := viper.GetString("mysql.DBName")

	path :=strings.Join([]string{user,":",passwd,"@tcp(",host,":",port,")/",dbName,"?charset=utf8&parseTime=true"},"")
	model.Database(path)
}

func LoadRabbitMQ()  {
	rabbitmq := viper.GetString("rabbitmq.RabbitMQ")
	user := viper.GetString("rabbitmq.RabbitMQUser")
	passwd := viper.GetString("rabbitmq.RabbitMQPassWord")
	host := viper.GetString("rabbitmq.RabbitMQHost")
	port := viper.GetString("rabbitmq.RabbitMQPort")


	path := strings.Join([]string{rabbitmq,"://",user,":",passwd,"@",host,":",port,"/"},"")
	model.RabbitMQ(path)
}
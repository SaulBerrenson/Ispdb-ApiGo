package main

import (
	configs "MailConfigHandler/Configs"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
)

var dbContext *configs.ConfigDB
var router *mux.Router


type ApiConfig struct {
	Domain *string
	Info *configs.EmailProvider

}

func contains(s []ApiConfig, predict string) bool {
	for _,domain:= range s {
		if *domain.Domain == predict {return true}
	}

	return false
}

func main() {

	countThreadsForUpdater := flag.Int("update-threads", 100, "an int")
	portApi := flag.Int("port", 37896, "an int")
	needUpdate := flag.Bool("update-cache", false, "a bool")
	flag.Parse()

	var config = configs.CreateConfigDB()
	config, _ = config.Update(needUpdate, countThreadsForUpdater)
	config.UpdateCache()
	fmt.Println("Найдено: ", config.GetCount())
	config, _ = config.PrepareDict()
	dbContext = config

	InitApi(*portApi)
}

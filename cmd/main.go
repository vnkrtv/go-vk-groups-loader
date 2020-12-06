package main

import (
	"log"
	"time"

	"github.com/vnkrtv/go-vk-groups-loader/pkg/service"
)

const (
	cfgPath    = "config/config.json"
	groupsPath = "config/groups.json"
)

func main() {
	cfg, err := service.GetConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	newsService, err := service.NewNewsService(
		cfg.VKToken, cfg.PGUser, cfg.PGPass, cfg.PGHost, cfg.PGPort, cfg.PGName)
	if err != nil {
		log.Fatal(err)
	}

	groupsScreenNames,err := service.GetGroupsScreenNames(groupsPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := newsService.InitDB(); err != nil {
		log.Println(err)
	}
	if err := newsService.AddNewsSources(groupsScreenNames); err != nil {
		log.Fatal(err)
	}
	for {
		if err := newsService.LoadNews(100); err != nil {
			log.Println(err)
		} else {
			log.Println()
		}
		time.Sleep(time.Duration(cfg.Interval) * time.Second)
	}
}

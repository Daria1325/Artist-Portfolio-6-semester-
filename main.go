package main

import (
	"fmt"
	cnfg "github.com/daria/Portfolio/backend/config"
	"github.com/daria/Portfolio/backend/database"
	"github.com/daria/Portfolio/backend/server"
)

var (
	ConfigPath = "data/configs/dataConfig.toml"
)

func main() {
	config, err := cnfg.NewConfigPath(ConfigPath)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}
	server.MainServer.Repo = database.Init(config)
	defer server.MainServer.Repo.Close()
	server.MainServer.StatusUser = false

	err = server.Start(config)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}
}

package main

import (
	"fmt"
	"github.com/daria/Portfolio/backend/cmd/client"
	"github.com/daria/Portfolio/backend/cmd/database"
	"github.com/daria/Portfolio/backend/cmd/server"
	cnfg "github.com/daria/Portfolio/backend/config"
)

var (
	ConfigPath = "configs/dataConfig.toml"
)

func configService() (*cnfg.Config, error) {
	config, err := cnfg.NewConfigPath(configPath)
	if err != nil {
		return nil, err
	}
	err = client.Connect.InitConn(config.BindAddrHost, config.BindAddrServer, config.FileSize)
	if err != nil {
		return nil, err
	}
	return config, nil
}
func main() {
	config, err := cnfg.NewConfigPath(ConfigPath)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}
	MainServer := server.Server{}
	MainServer.Repo = database.Init(config)
	defer MainServer.Repo.Close()

	err = client.Start(config)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}
	defer client.Connect.Conn.Close()
}

package utils

import (
	"github.com/gookit/ini/v2"
	"github.com/sandata/cops/starcoin"
	"github.com/urfave/cli/v2"
)

func InitConfig (c *cli.Context) *starcoin.StarCoinConfigure{
	// LoadExists 将忽略不存在的文件
	config := c.String("config")
	if config=="" {
		config = "/etc/cops/config.ini"
	}
	err := ini.LoadExists(config, "not-exist.ini")
	if err != nil {
		panic(err)
	}
	networkType := ini.String("network_type")
	dataPath := ini.String("data_path")
	accessToken := ini.String("github_token")
	chainId := c.String("chain_id")
	password := c.String("password")
	projectName := c.String("project_name")
	network := c.String("network")
	starCoinConfigure := &starcoin.StarCoinConfigure{
		DataPath:    dataPath,
		ChainId:     chainId,
		Password:    password,
		NetworkType: networkType,
		Token:       accessToken,
		ProjectName: projectName,
		Network: network,
	}
	return starCoinConfigure
}

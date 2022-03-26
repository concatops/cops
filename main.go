package main

import (
    "fmt"
    "github.com/google/uuid"
    "github.com/sandata/cops/codelab"
    "github.com/sandata/cops/starcoin"
    "github.com/sandata/cops/utils"
    "github.com/urfave/cli/v2"
    "log"
    "os"
    "os/exec"
    "sort"
    "strings"
)

func main() {
    app := &cli.App{
    Name: "cops",
    Usage: "fight the loneliness!",
    Commands: []*cli.Command{
        {
            Name:        "accounts",
            Aliases:     []string{"acs"},
            Usage:       "list all blockchain account",
            Description: "This is show all account on blockchain",
            Action: func(c *cli.Context) error {
                coinConfigure :=utils.InitConfig(c)
                network:=coinConfigure.Network
                if network=="starcoin"{
                    accountList :=starcoin.AccountList(coinConfigure)
                    fmt.Printf("accountList is: %s\n", accountList)
                }
                return nil
            },
        },
        {
            Name:        "start",
            Aliases:     []string{"node"},
            Usage:       "start cops dev network",
            Description: "This is start a cops dev network",
            Action: func(c *cli.Context) error {
                coinConfigure :=utils.InitConfig(c)
                network:=coinConfigure.Network
                if network=="starcoin"{
                    uuid := uuid.New()
                    key := uuid.String()
                    cmd := exec.Command("bash", "-c", "starcoin -n "+
                        coinConfigure.NetworkType +" -d "+coinConfigure.DataPath+key)
                    stdout, err := cmd.Output()
                    if err != nil {
                        fmt.Printf("error: %+v\n", err)
                    }
                    fmt.Printf("The blockchain id is %s\n", stdout)
                    starcoin.GetCoin(coinConfigure)
                }
                return nil
            },
        },
        {
            Name:        "init",
            Aliases:     []string{"i"},
            Usage:       "init cops  example contract project",
            Description: "This is init a  example github project",
            Action: func(c *cli.Context) error {
                coinConfigure :=utils.InitConfig(c)
                //init cops example project
                network:=coinConfigure.Network
                if network=="starcoin"{
                    codelab.CreateTemplateRepo(coinConfigure.Token,
                        coinConfigure.ProjectName,true)
                }
                return nil
            },
        },
        {
            Name:        "deploy",
            Aliases:     []string{"d"},
            Usage:       "deploy cops  example contract project",
            Description: "This is deploy a  example github project",
            Action: func(c *cli.Context) error {
                //deploy cops example project
                coinConfigure :=utils.InitConfig(c)
                network:=coinConfigure.Network
                if network=="starcoin"{
                    if coinConfigure.UserAddress=="" {
                        //create new account
                        if strings.Compare(coinConfigure.NetworkType,"dev")!=0{
                            coinConfigure.ChainId=""
                        }
                        userAddress :=starcoin.CreateUser(coinConfigure)
                        coinConfigure.UserAddress = userAddress
                    }
                    starcoin.BuildContract(coinConfigure)
                }
                return nil
            },
        },
    },

    Action: func(c *cli.Context) error {
      return nil
    },

    Flags: []cli.Flag{
          &cli.StringFlag{
            Name:    "access_token",
            Aliases: []string{"t"},
            Usage:   "Set GitHub access_token",
          },
          &cli.StringFlag{
            Name:    "config",
            Aliases: []string{"f"},
            Usage:   "Set cops config path",
          },
          &cli.StringFlag{
            Name:    "username",
            Aliases: []string{"u"},
            Usage:   "Set github username",
          },
          &cli.StringFlag{
            Name:    "password",
            Aliases: []string{"pwd"},
            Usage:   "Set blockchain password",
          },
          &cli.StringFlag{
            Name:    "user_address",
            Aliases: []string{"uar"},
            Usage:   "Set blockchain user_address",
          },
          &cli.StringFlag{
            Name:    "project_name",
            Aliases: []string{"pn"},
            Usage:   "Set project_name",
          },
          &cli.StringFlag{
             Name:    "chain_id",
             Aliases: []string{"cid"},
             Usage:   "Set blockchain chain_id",
          },
          &cli.StringFlag{
            Name:    "network",
            Aliases: []string{"n"},
            Usage:   "Set which blockchain network you want ",
          },
          &cli.StringFlag{
             Name:    "private_flag",
             Aliases: []string{"pf"},
             Usage:   "Set github project private flag",
          },
     },
  }
  sort.Sort(cli.FlagsByName(app.Flags))
  sort.Sort(cli.CommandsByName(app.Commands))
  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
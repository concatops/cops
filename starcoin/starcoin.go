package starcoin

import (
	"fmt"
	"os/exec"
)

type StarCoinConfigure struct {
	DataPath    string
	ChainId     string
	NetworkType string
	Password    string
	UserAddress string
	Token       string
	ProjectName string
	Network     string
}

func CreateUser(starCoinConfigure *StarCoinConfigure) string{
	dataPath := starCoinConfigure.DataPath
	chainId := starCoinConfigure.ChainId
	networkType := starCoinConfigure.NetworkType
	password := starCoinConfigure.Password
	cmd := exec.Command("bash", "-c", "starcoin -c "+
		dataPath+chainId+"/"+networkType+"/starcoin.ipc account create -p "+password+
		" | sed 's/,/\\n/g' | grep \"address\" | sed 's/:/\\n/g' | sed '1d' | sed 's/}//g' | sed $'s/\\\"//g' | sed 's/^[ \\t]*//g'")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: %+v\n", err)
	}
	fmt.Printf("The user address is %s\n", stdout)
	return string(stdout)
}

func changeToml(starCoinConfigure *StarCoinConfigure){
	userAddress :=starCoinConfigure.UserAddress
	cmd := exec.Command("bash", "-c", "sed -i 's/$uaddress/"+userAddress+"/g' Move.toml")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: %+v\n", err)
	}
	fmt.Printf("changeToml is %s\n", stdout)
}


func GetCoin(starCoinConfigure *StarCoinConfigure){
	dataPath := starCoinConfigure.DataPath
	chainId := starCoinConfigure.ChainId
	networkType := starCoinConfigure.NetworkType
	password := starCoinConfigure.Password
	userAddress :=starCoinConfigure.UserAddress
	getCoinCmd := dataPath+chainId+"/"+networkType+"/starcoin.ipc dev get-coin -v 10000000STC"
	if userAddress !=""{
		getCoinCmd = getCoinCmd +" -p "+password+"  "+userAddress
	}
	cmd := exec.Command("bash", "-c", "starcoin -c "+getCoinCmd)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf(getCoinCmd+"\n")
		fmt.Printf("error: %+v\n", err)
	}
	fmt.Printf("Get Coin end %s\n", stdout)
}

func unlockAccount(starCoinConfigure *StarCoinConfigure){
	dataPath := starCoinConfigure.DataPath
	chainId := starCoinConfigure.ChainId
	networkType := starCoinConfigure.NetworkType
	password := starCoinConfigure.Password
	userAddress :=starCoinConfigure.UserAddress
	cmd := exec.Command("bash", "-c", "starcoin -c "+
		dataPath+chainId+"/"+networkType+"/starcoin.ipc account unlock -p "+password+"  "+userAddress)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: %+v\n", err)
	}
	fmt.Printf("The user address is %s\n", stdout)
}

func AccountList(starCoinConfigure *StarCoinConfigure) string{
	dataPath := starCoinConfigure.DataPath
	chainId := starCoinConfigure.ChainId
	networkType := starCoinConfigure.NetworkType
	password := starCoinConfigure.Password
	userAddress :=starCoinConfigure.UserAddress
	baseCommand :="starcoin -c "+dataPath+chainId+"/"+networkType+"/starcoin.ipc account list"
	if password!=""{
		baseCommand = baseCommand+" -p "+password
	}
	if userAddress!=""{
		baseCommand = baseCommand+" "+userAddress
	}
	cmd := exec.Command("bash", "-c", baseCommand)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: %+v\n", err)
	}
	return string(stdout)
}

func deployContract(starCoinConfigure *StarCoinConfigure){
	dataPath := starCoinConfigure.DataPath
	chainId := starCoinConfigure.ChainId
	networkType := starCoinConfigure.NetworkType
	password := starCoinConfigure.Password
	userAddress :=starCoinConfigure.UserAddress
	cmd := exec.Command("bash", "-c", "starcoin -c "+
		dataPath+chainId+"/"+networkType+"/starcoin.ipc dev deploy -p "+password+"  -b "+userAddress)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: %+v\n", err)
	}
	fmt.Printf("The user contract is %s\n", stdout)
}

func BuildContract(starCoinConfigure *StarCoinConfigure){
	changeToml(starCoinConfigure)
	cmd := exec.Command("bash", "-c", "mpm release")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: %+v\n", err)
	}
	fmt.Printf("The date is %s\n", stdout)
	unlockAccount(starCoinConfigure)
	deployContract(starCoinConfigure)
}

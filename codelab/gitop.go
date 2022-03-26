package codelab

import (
	"fmt"
	//"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
  "github.com/google/go-github/v43/github"
	"log"
	"runtime"
	"os/exec"
	"net/http"
	"io/ioutil"
)

func CreateRepo(accessToken string,name string, projectType bool) bool {
    ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	// list all repositories for the authenticated user
	// create a new private repository named "foo"
    repo := &github.Repository{
    	Name:    &name,
    	Private: &projectType,
    }
	_, _, error:= client.Repositories.Create(ctx, "", repo)
	if error != nil {
		log.Fatal(error)
		return false
	}else{
		return true
	}
}


func CreateTemplateRepo(accessToken string,name string, projectType bool) bool {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	// list all repositories for the authenticated user
	// create a new private repository named "foo"
	repo := &github.TemplateRepoRequest{
		Name:    &name,
		Private: &projectType,
	}
	_, _, error:= client.Repositories.CreateFromTemplate(ctx, "concatops","template", repo)
	if error != nil {
		log.Fatal(error)
		return false
	}else{
		return true
	}
}

func DownLoadExample (){
	// Get the data
	resp, err := http.Get("http://116.62.185.67:8000/contract_example.zip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("contract_example.zip", data, 0644)
}

/** Init Project **/
func InitProject (projectName string){
	sysType := runtime.GOOS
	if sysType == "linux" {
		cmd := exec.Command("bash", "-c", "git init")
		output, err := cmd.CombinedOutput()
		fmt.Printf("output:\n%s\n", output)
		if err != nil {
			fmt.Printf("error: %+v\n", err)
		}
	}
	//if sysType == "windows" {
	//	// windows系统
	//	cmd := exec.Command("cmd.exe")
	//	//project_path:= folderPath+"\\"+projectName
	//	cmdExec := `git init `
	//	fmt.Println(fmt.Sprintf(`/c %s`, cmdExec))
	//	//核心点,直接修改执行命令方式
	//	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(`/c %s`, cmdExec), HideWindow: true}
	//	output, err := cmd.Output()
	//	fmt.Printf("output:\n%s\n", output)
	//	if err != nil {
	//		fmt.Printf("error: %+v\n", err)
	//	}
	//}
}
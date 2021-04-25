package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)
import "flag"
import "net/http"

const (
	user = "yuanzhe"
	auth = "dlgamelog"
)
var date = flag.String("t","","-t 2021_04_25",)
func main(){
	flag.Parse()
	url := fmt.Sprintf("http://115.231.79.188:12601/v1/sj/?date=%s&user=%s&auth=%s",*date,user,auth)
	urlDL := fmt.Sprintf("http://115.231.79.188:12601/v1/sj/download?user=%s&auth=%s",user,auth)
	client := &http.Client{}
	req,_ := http.NewRequest("GET",url,nil)
	req.Header.Set("accept","application/json")
	response,err := client.Do(req)
	if err != nil{
		fmt.Println(" No connection could be made because the target machine actively refused it")
		return
	}
	defer response.Body.Close()
	resByte,_ := ioutil.ReadAll(response.Body)
	if string(resByte) != "\"ok\""{
		fmt.Println(string(resByte))
		return
	}
	reqDownload,_ := http.NewRequest("GET",urlDL,nil)
	reqDownload.Header.Set("accept","application/json")
	responseDL,err := client.Do(reqDownload)
	if err != nil{
		fmt.Println("No connection could be made because the target machine actively refused it")
	}
	defer responseDL.Body.Close()
	file, err := os.Create("sgsjlog.tgz")
	if err !=nil{
		panic(err)
	}
	_,err = io.Copy(file,responseDL.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("download log file to sgsjlog.tgz")
	s:= exec.Command("rm","rm","")
	s.Run()
}
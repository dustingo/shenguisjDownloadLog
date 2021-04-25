package models

import (
	"errors"
	"os/exec"
)

type Admin struct {
	User string
	Auth string
}

var yuanzhe Admin = Admin{
	User: "yuanzhe",
	Auth: "dlgamelog",
}

func DownloadGameLog(date string) error {
	cmd := exec.Command("dlog", "-t", date)
	err := cmd.Run()
	if err != nil {
		return errors.New("download error")
	}
	cmdCompress := exec.Command("tar", "zcvf", "sgsjlog.tgz", "/export/hefu/logs")
	err = cmdCompress.Run()
	if err != nil {
		return errors.New("error while compressing log directory")
	}
	return nil
}
func (A *Admin) GetUser() (user string) {
	return yuanzhe.User
}
func (A *Admin) GetAuth() (auth string) {
	return yuanzhe.Auth
}
func DownloadComprocessFile() string {
	return "./sgsjlog.tgz"
}

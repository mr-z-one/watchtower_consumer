package handler

import (
	"errors"
	"watchtower_consumer/Constants"
	custom_color "watchtower_consumer/custom_Color"
	"watchtower_consumer/tools/Interface"
	"watchtower_consumer/tools/program"
)

var procedures = make(map[string]Interface.IExecutable)

func Get_Program(cmd string) (Interface.IExecutable, error) {
	exe, exist := procedures[cmd]

	if !exist {
		custom_color.Error()("'" + cmd + "'" + " " + "not Found\n")
		return nil, errors.New("'" + cmd + "'" + " " + "not Found")
	}
	return exe, nil
}

func Register_All() {
	register(Constants.APP_SUBFINDER, &program.SubFinder{Name: Constants.APP_SUBFINDER, Use_proxy: true})

}

func register(cmd string, procedure Interface.IExecutable) {
	_, exist := procedures[cmd]

	if exist {
		custom_color.Error()("'" + cmd + "'" + " " + "already registered\n")
		panic("")
	}

	procedures[cmd] = procedure
}

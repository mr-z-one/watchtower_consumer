package env

import (
	"os"
	"watchtower_consumer/Constants"
	"watchtower_consumer/config"
	custom_color "watchtower_consumer/custom_Color"
)

func Init() {

	conf := config.ReadeConfig("config.json")

	err := os.Setenv(Constants.PROXY_KEY, conf.GetProxy_url())

	if err != nil {
		custom_color.Error()("\n[-] %v\n", err)
		panic("")

	}

	err = os.Setenv(Constants.BASE_PATH_KEY, conf.Base_Path_App)
	if err != nil {
		custom_color.Error()("\n[-] %v\n", err)
		panic("")

	}
	custom_color.Succeed()("[+] initialization of env, successes\n")
}

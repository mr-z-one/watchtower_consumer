package env

import (
	"os"
	"watchtower_consumer/Constants"
	"watchtower_consumer/config"
	custom_color "watchtower_consumer/custom_Color"
	"watchtower_consumer/utils"
)

func Init() {

	conf := config.ReadeConfig("config.json")

	err := os.Setenv(Constants.PROXY_KEY, conf.GetProxy_url())

	utils.FailOnErrorPanic(err, "", nil)

	err = os.Setenv(Constants.BASE_PATH_KEY, conf.Base_Path_App)
	utils.FailOnErrorPanic(err, "", nil)

	custom_color.Succeed()("[+] initialization of env, successes\n")
}

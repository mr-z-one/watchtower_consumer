package tools

import (
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"watchtower_consumer/Constants"
	custom_color "watchtower_consumer/custom_Color"
	"watchtower_consumer/utils"
)

func Check_Proxy() bool {

	proxyUrl, _ := url.Parse(os.Getenv(Constants.PROXY_KEY))
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	_, err := myClient.Head("https://www.amd.com/")

	e := utils.FailOnError(err, "", nil)
	if e {
		return false
	}

	return true

}

func RunApp(name string, args ...string) ([]byte, error) {
	base := os.Getenv(Constants.BASE_PATH_KEY)
	if base == "" {
		custom_color.Error()("[-] base path not Found in env \n")
		panic("")
	}

	full_path := base + name

	if _, err := os.Stat(full_path); err != nil {
		custom_color.Error()("[-] App not Found in %s \n", full_path)
		panic("")
	}

	return exec.Command(name, args...).Output()
}

package config

import (
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strconv"
	custom_color "watchtower_consumer/custom_Color"
)

type user_config struct {
	Port          int    `json:"port"`
	Proxy         string `json:"proxy"`
	Base_Path_App string `json:"base_path_app"`
}

func (conf *user_config) GetProxy_url() string {
	return conf.Proxy + ":" + strconv.Itoa(conf.Port)
}

func ReadeConfig(name string) *user_config {

	r := regexp.MustCompile(`[^.]+\.json$`)
	custom_color.Succeed()("[+] start reading config.json\n")
	if !r.MatchString(name) {
		custom_color.Error()("[-] this not a json file !!\n")

		return nil
	}
	config, err := os.Open(name)

	if err != nil {
		custom_color.Error()("[-] %v\n", err)
		return nil
	}
	defer config.Close()
	data, err := io.ReadAll(config)
	if err != nil {
		custom_color.Error()("[-] %v\n", err)
		return nil
	}

	var userConfig user_config

	err = json.Unmarshal(data, &userConfig)

	if err != nil {
		custom_color.Error()("[-] %v\n", err)
		return nil
	}

	return &userConfig
}

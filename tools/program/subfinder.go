package program

import (
	"errors"
	"os"
	"watchtower_consumer/Constants"
	custom_color "watchtower_consumer/custom_Color"
	"watchtower_consumer/tools"
)

type SubFinder struct {
	Name      string
	Use_proxy bool
}

func (s *SubFinder) Execute(args ...string) ([]byte, error) {
	if len(args) == 0 {
		custom_color.Error()("[-] domain required  \n")
		return nil, errors.New("[-] domain required ")
	}
	app_args := []string{}
	app_name := s.Name
	domain := args[0]

	app_args = append(app_args, "-all")
	app_args = append(app_args, "-silent")
	if s.Use_proxy {
		proxy := os.Getenv(Constants.PROXY_KEY)
		if proxy == "" {
			return nil, errors.New("[-] proxy  not Found in env")
		}

		if !tools.Check_Proxy() {
			return nil, errors.New("[-] proxy Offline")
		}

		app_args = append(app_args, "-proxy")
		app_args = append(app_args, proxy)
	}

	app_args = append(app_args, "-d")
	app_args = append(app_args, domain)

	return tools.RunApp(app_name, app_args...)

}

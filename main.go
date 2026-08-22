package main

import (
	"fmt"
	"watchtower_consumer/Constants"
	custom_color "watchtower_consumer/custom_Color"
	"watchtower_consumer/env"
	"watchtower_consumer/tools"
	"watchtower_consumer/tools/handler"
)

func main() {

	env.Init()
	handler.Register_All()

	if tools.Check_Proxy() {
		custom_color.Succeed()("[+] Proxy Turn on\n")
	} else {
		custom_color.Warning()("[!] Proxy Turn off\n")
	}

	ch := make(chan string)

	go func(ch chan string) {

		exe, _ := handler.Get_Program(Constants.APP_SUBFINDER)
		data, _ := exe.Execute("sku.ac.ir")
		ch <- string(data)
	}(ch)

	fmt.Println("app end..")

	fmt.Println(<-ch)
	// r := exec.Command("curl", "-x", "http://127.0.0.1:10808", "-s", "-I", "http://www.google.com")
	// b, err := r.Output()
	// fmt.Println(err)
	// fmt.Println(string(b))
	// tools.exe

}

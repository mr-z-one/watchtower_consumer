package utils

import (
	custom_color "watchtower_consumer/custom_Color"
)

func FailOnError(err error, msg string, fn func()) bool {

	if err != nil {
		custom_color.Error()(msg)
		custom_color.Error()("\n[-] %v\n", err)
		if fn != nil {

			fn()
		}
		return true
	}
	return false
}

func FailOnErrorPanic(err error, msg string, fn func()) {

	if err != nil {
		custom_color.Error()(msg)
		custom_color.Error()("\n[-] %v\n", err)
		if fn != nil {

			fn()
		}
		panic("")

	}
}

package custom_color

import "github.com/fatih/color"

func Warning() func(format string, a ...interface{}) {

	return color.New(color.FgYellow, color.Bold).PrintfFunc()

}
func Reset() func(format string, a ...interface{}) {

	return color.New(color.ResetBlinking).PrintfFunc()

}

func Succeed() func(format string, a ...interface{}) {

	return color.New(color.FgGreen, color.Bold).PrintfFunc()

}

func Error() func(format string, a ...interface{}) {

	return color.New(color.FgRed, color.Bold).PrintfFunc()

}

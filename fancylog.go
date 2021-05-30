package flog

import "fmt"

func LogSln(str string) {
	fmt.Println(SuccessSym, SuccessColor(str))
}

func LogSf(format string, a ...interface{}) {
	fmt.Printf("%s %s\n", SuccessSym, SuccessColor(format, a))
}

func LogWln(str string) {
	fmt.Println(WarningSym, WarnColor(str))
}

func LogWf(format string, a ...interface{}) {
	fmt.Printf("%s %s\n", WarningSym, WarnColor(format, a))
}

func LogEln(str string) {
	fmt.Println(ErrorSym, ErrorColor(str))
}

func LogEf(format string, a ...interface{}) {
	fmt.Printf("%s %s\n", ErrorSym, ErrorColor(format, a))
}

func Logln(symbol Symbol, str string) {
	fmt.Println(symbol, str)
}

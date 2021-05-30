package main

import (
	"fmt"

	"fancylog"
)

func main() {
	fmt.Println(fancylog.Success)
	fmt.Println(fancylog.Warning)
	fmt.Println(fancylog.Error)
	fmt.Println(fancylog.Info)
	fmt.Println(fancylog.Symbol('\u2764'))
}

package main

import (
	"fmt"

	"github.com/calvernaz/flog"
	"gopkg.in/gookit/color.v1"
)

func main() {
	// just symbols
	fmt.Println(flog.SuccessSym)
	fmt.Println(flog.WarningSym)
	fmt.Println(flog.ErrorSym)
	fmt.Println(flog.InfoSym)

	// add custom symbol not in the library
	symbol := flog.Symbol{
		Unicode: '\u2764', // ❤
		Color:   color.FgRed.Sprintf,
	}
	fmt.Println(symbol)
}

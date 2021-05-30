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
	fmt.Println(flog.StopSym)

	// add custom symbol not in the library
	symbol := flog.Symbol{
		Unicode: '\u2764', // ❤
		Color:   color.FgRed.Sprintf,
	}
	fmt.Println(symbol)

	flog.LogSln("Successfully created!")
	flog.LogSf("Image %s created in the background\n", "cat.png")

	flog.LogWln("Image size too large!")
	flog.LogWf("Image %s should have maximum size of 48x48\n", "cat.png")

	flog.LogEln("Failed to create!")
	flog.LogEf("Image %s failed to create with size 64x64", "cat.png")
}

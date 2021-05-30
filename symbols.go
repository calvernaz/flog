package flog

import (
	"fmt"

	"gopkg.in/gookit/color.v1"
)



// Symbol type
type Symbol struct {
	// https://unicode-table.com
	Unicode rune
	Color   Color
}

var (
	InfoSym    = Symbol{'\u2139', Color(color.Info.Sprintf)}    // ℹ
	SuccessSym = Symbol{'\u2714', Color(color.Success.Sprintf)} // ✓
	WarningSym = Symbol{'\u26A0', Color(color.Warn.Sprintf)}    // ⚠
	ErrorSym   = Symbol{'\u2717', Color(color.Error.Sprintf)}   // ✗
)

func (s Symbol) String() string {
	return fmt.Sprintf("%s", s.Color("%c", s.Unicode))
}

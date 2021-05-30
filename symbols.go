package flog

import (
	"fmt"
)

// Symbol type
type Symbol struct {
	// https://unicode-table.com
	Unicode rune
	Color   Color
}

var (
	InfoSym    = Symbol{'\u2139', InfoColor}    // ℹ
	SuccessSym = Symbol{'\u2714', SuccessColor} // ✓
	WarningSym = Symbol{'\u26A0', WarnColor}    // ⚠
	ErrorSym   = Symbol{'\u2717', ErrorColor}   // ✗
	StopSym = Symbol{'\u2B23', StopColor} // ⬣
)

func (s Symbol) String() string {
	return fmt.Sprintf("%s", s.Color("%c", s.Unicode))
}

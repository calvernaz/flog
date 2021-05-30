package fancylog

import "fmt"

// Symbol https://unicode-table.com
type Symbol rune

const (
	Info    Symbol = '\u2139' // ℹ
	Success Symbol = '\u2714' // ✓
	Warning Symbol = '\u26A0' // ⚠
	Error   Symbol = '\u2717' // ✗
)

func (s Symbol) String() string {
	return fmt.Sprintf("%c", s)
}

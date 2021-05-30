package flog

import "gopkg.in/gookit/color.v1"

type Color func(format string, a ...interface{}) string

var (
	InfoColor = Color(color.Info.Sprintf)
	SuccessColor = Color(color.Success.Sprintf)
	WarnColor = Color(color.Warn.Sprintf)
	ErrorColor = Color(color.Error.Sprintf)
)

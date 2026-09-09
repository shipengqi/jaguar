package ui

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

var (
	ColorBlue   = lipgloss.Color("#7571F9")
	ColorGreen  = lipgloss.Color("#02BF87")
	ColorYellow = lipgloss.Color("#FFE600")
	ColorRed    = lipgloss.Color("#ED567A")
	ColorGray   = lipgloss.Color("#B6B4AD")
)

type StringStyle struct {
	IsBold bool
	Color  color.Color
}

type FrameStyle struct {
	Padding []int
	Color   color.Color
}

func MakeStyledString(s string, style *StringStyle) string {
	if style == nil {
		return s
	}
	return lipgloss.NewStyle().Foreground(style.Color).Bold(style.IsBold).Render(s)
}

func MakeStyledFrame(s string, style *FrameStyle) string {
	if style == nil {
		return s
	}
	return lipgloss.NewStyle().
		Padding(style.Padding...).
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderLeftForeground(style.Color).
		Render(s)
}

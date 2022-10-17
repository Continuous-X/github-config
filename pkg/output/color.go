package output

import "github.com/fatih/color"

func getRed() *color.Color {
	return color.New(color.FgRed).Add(color.BlinkSlow)
}

func getGreen() *color.Color {
	return color.New(color.FgGreen).Add(color.Underline)
}

func getYellow() *color.Color {
	return color.New(color.FgYellow)
}

func getWhite() *color.Color {
	return color.New(color.FgWhite)
}

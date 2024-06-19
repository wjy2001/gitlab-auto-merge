package main

import (
	"gitlab-auto-merge/gui/fonts"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

// Font return bundled font resource
func (*MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	return &fyne.StaticResource{
		StaticName:    "font.ttf",
		StaticContent: fonts.StkaitiTTf,
	}

}

func (*MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*MyTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

var outDir string = "./themes"

func main() {
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err = os.MkdirAll(outDir, 0755)
		if err != nil {
			panic(fmt.Sprintf("could not create directory: %s", err))
		}
	}

	themes := []Theme{
		{
			Name:          "Reui",
			Source:        "https://github.com/barrsan/reui-vscode-theme",
			Background:    "#282c34",
			Foreground:    "#f8f8f2",
			Black:         "#21222c",
			Blue:          "#bd93f9",
			Cyan:          "#8be9fd",
			Green:         "#8dc891",
			Magenta:       "#ff79c6",
			Red:           "#ff5555",
			White:         "#f8f8f2",
			Yellow:        "#f1fa8c",
			BrightBlack:   "#6272a4",
			BrightBlue:    "#d6acff",
			BrightCyan:    "#a4ffff",
			BrightGreen:   "#8dc891",
			BrightMagenta: "#ff92df",
			BrightRed:     "#ff6e6e",
			BrightWhite:   "#ffffff",
			BrightYellow:  "#ffffa5",
		},
	}

	terminals := []Terminal{
		&Alacritty{fileExtension: "yml"},
		&Kitty{fileExtension: "conf"},
	}

	for _, theme := range themes {
		for _, terminal := range terminals {
			err := terminal.WriteFile(theme)
			if err != nil {
				fmt.Printf("failed to write theme %s for terminal %s\n", theme.Name, reflect.TypeOf(terminal).String())
			}
		}
	}
}

type Theme struct {
	Name          string
	Source        string // link to repository or webpage
	Background    string
	Foreground    string
	Black         string
	Blue          string
	Cyan          string
	Green         string
	Magenta       string
	Red           string
	White         string
	Yellow        string
	BrightBlack   string
	BrightBlue    string
	BrightCyan    string
	BrightGreen   string
	BrightMagenta string
	BrightRed     string
	BrightWhite   string
	BrightYellow  string
}

type Terminal interface {
	WriteFile(theme Theme) error
}

type Alacritty struct {
	fileExtension string
}

func (a *Alacritty) WriteFile(theme Theme) error {
	f, err := os.Create(fmt.Sprintf("%s/%s.%s", outDir, theme.Name, a.fileExtension))
	if err != nil {
		return err
	}

	t :=
		"# " + theme.Name + " colors" + "\n" +
			"# " + theme.Source + "\n" +
			"colors:" + "\n" +
			"  # Default colors" + "\n" +
			"  primary:" + "\n" +
			"    background: " + strings.Replace(theme.Background, "#", "0x", 1) + "\n" +
			"    foreground: " + strings.Replace(theme.Foreground, "#", "0x", 1) + "\n" +
			"\n" +
			"  # Normal colors" + "\n" +
			"  normal:" + "\n" +
			"    black: " + strings.Replace(theme.Black, "#", "0x", 1) + "\n" +
			"    red: " + strings.Replace(theme.Red, "#", "0x", 1) + "\n" +
			"    green:" + strings.Replace(theme.Green, "#", "0x", 1) + "\n" +
			"    yellow: " + strings.Replace(theme.Yellow, "#", "0x", 1) + "\n" +
			"    blue: " + strings.Replace(theme.Blue, "#", "0x", 1) + "\n" +
			"    magenta: " + strings.Replace(theme.Magenta, "#", "0x", 1) + "\n" +
			"    cyan: " + strings.Replace(theme.Cyan, "#", "0x", 1) + "\n" +
			"    white: " + strings.Replace(theme.White, "#", "0x", 1) + "\n" +
			"\n" +
			"  # Bright colors" + "\n" +
			"  bright:" + "\n" +
			"    black: " + strings.Replace(theme.BrightBlack, "#", "0x", 1) + "\n" +
			"    red: " + strings.Replace(theme.BrightRed, "#", "0x", 1) + "\n" +
			"    green: " + strings.Replace(theme.BrightGreen, "#", "0x", 1) + "\n" +
			"    yellow: " + strings.Replace(theme.BrightYellow, "#", "0x", 1) + "\n" +
			"    blue: " + strings.Replace(theme.BrightBlue, "#", "0x", 1) + "\n" +
			"    magenta: " + strings.Replace(theme.BrightMagenta, "#", "0x", 1) + "\n" +
			"    cyan: " + strings.Replace(theme.BrightCyan, "#", "0x", 1) + "\n" +
			"    white: " + strings.Replace(theme.BrightWhite, "#", "0x", 1) + "\n"

	_, err = f.WriteString(t)

	return err
}

type Kitty struct {
	fileExtension string
}

func (k *Kitty) WriteFile(theme Theme) error {
	f, err := os.Create(fmt.Sprintf("%s/%s.%s", outDir, theme.Name, k.fileExtension))
	if err != nil {
		return err
	}

	t :=
		"# vim:ft=kitty" + "\n" +
			"\n" +
			"## name: " + theme.Name + "\n" +
			"## source: " + theme.Source + "\n" +
			"\n" +
			"foreground           " + theme.Foreground + "\n" +
			"background           " + theme.Background + "\n" +
			"selection_foreground " + theme.Foreground + "\n" +
			"selection_background " + theme.Background + "\n" +
			"\n" +
			"# black" + "\n" +
			"color0      " + theme.Black + "\n" +
			"color8      " + theme.BrightBlack + "\n" +
			"\n" +
			"# red" + "\n" +
			"color1      " + theme.Red + "\n" +
			"color9      " + theme.BrightRed + "\n" +
			"\n" +
			"# green" + "\n" +
			"color2      " + theme.Green + "\n" +
			"color10     " + theme.BrightGreen + "\n" +
			"\n" +
			"# yellow" + "\n" +
			"color3      " + theme.Yellow + "\n" +
			"color11     " + theme.BrightYellow + "\n" +
			"\n" +
			"# blue" + "\n" +
			"color4      " + theme.Blue + "\n" +
			"color12     " + theme.BrightBlue + "\n" +
			"\n" +
			"# magenta" + "\n" +
			"color5      " + theme.Magenta + "\n" +
			"color13     " + theme.BrightMagenta + "\n" +
			"\n" +
			"# cyan" + "\n" +
			"color6      " + theme.Cyan + "\n" +
			"color14     " + theme.BrightCyan + "\n" +
			"\n" +
			"# white" + "\n" +
			"color7      " + theme.White + "\n" +
			"color15     " + theme.BrightWhite + "\n"

	_, err = f.WriteString(t)
	return err
}

package main

import (
	"colors/folder"
	"fmt"
	"os"
	"strings"
)

func main() {
	colorName := os.Args[1]
	letter := os.Args[2]
	word := os.Args[3]
	colorName = strings.ToLower(colorName)
	res := folder.ColorPicker(colorName)
	ansiColor := folder.ANSIColor(res.R, res.G, res.B)
	letterFields := strings.Split(letter, "")
	for _, char := range word {
		found := false
		for _, val := range letterFields {
			if string(char) == string(val) {
				found = true
				fmt.Printf("%s%s\x1b[0m", ansiColor,val)
				continue
			}
		}
		if !found{
			fmt.Print(string(char))

		}
		
	}
	fmt.Println()
	
}
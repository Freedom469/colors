package main

import (
	"colors/utils"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func GetStartingPoint(ascii int) int {
	return (ascii-32)*9 + 1
}

func preprocessArgs() {
	for i, arg := range os.Args {
		if !strings.HasPrefix(arg, "-") && strings.Contains(arg, "=") {
			os.Args[i] = "--" + arg
		}
	}
}

func main() {
	filePath := "standard.txt"
	preprocessArgs()
	args := os.Args[1:]

	if len(args) < 2 || len(args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> something")
		os.Exit(0)
	}

	pattern := `^--color=.+(-h .+)?(-i .+)$`

	regex := regexp.MustCompile(pattern)

	// Join the arguments into a single string with spaces between them
	argumentString := strings.Join(args, " ")
	// fmt.Println(argumentString)

	// Match the argument string against the regex pattern
	if !regex.MatchString(argumentString) {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> something")
		os.Exit(0)
	}

	var input string
	var Highlight string
	var color string

	isFlag := false

	flag.StringVar(&input, "i", "", "Gets The User Input")
	flag.StringVar(&Highlight, "h", "", "the letter or letters that you can chose to be colored.")
	flag.StringVar(&color, "color", "", "the color desired by the user")

	flag.Parse()
	if color == ""{
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=< CAN'T BE EMPTY > <letters to be colored> something")
		os.Exit(0)
	}

	if Highlight != "" {
		isFlag = true
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error Opening File: %v", err)
	}

	fileContent := strings.Split(string(file), "\n")
	ansiColor := utils.GetColor(color)
	inputString := strings.Split(input, "\\n")
	for _, part := range inputString {
		for i := 0; i < 8; i++ {
			for _, char := range part {
				index := GetStartingPoint(int(char))
				line := fileContent[index+i]
				if isFlag {
					if strings.ContainsRune(Highlight, char) {
						fmt.Printf("%s%s\x1b[0m", ansiColor, line)
					} else {
						fmt.Print(line)
					}
				}
				if !isFlag {
					fmt.Printf("%s%s\x1b[0m", ansiColor, line)
				}
			}
			fmt.Println()
		}
	}
}

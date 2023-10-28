package main

import (
	"fmt"
	"os"
	str "strings"

	gk "github.com/jwalton/gchalk"
)

var inputFile string = "."
var targetOs string = os.Getenv("GOOS")
var targetArch string = os.Getenv("GOARCH")
var out string = ""

func PrintUsage() {
	fmt.Println("Usage: gob --option=<option> <file (optional)>")
	fmt.Println("Use gob --help for a list of commands")
}

func ParseArgs() {
	for _, arg := range os.Args[1:] {
		if str.HasPrefix(arg, "--") {
			//get only the option name
			option := arg[:str.Index(arg, "=")][2:]
			value := arg[str.Index(arg, "=")+1:]
			switch option {
			case "help":
				fmt.Println(`Commands:
	--help: show this help
	--os: set target platform OS
	--arch: set target platform architecture
	--out: set output file
	Example:
	gob --os=linux --arch=amd64 --out=linux_amd64 main.go`)
			case "os":
				targetOs = value
			case "arch":
				targetArch = value
			case "out":
				out = value
			default:
				PrintUsage()
			}
			continue
		} else {
			if inputFile == "." {
				inputFile = arg
			} else {
				fmt.Println(gk.Red("Error: ") + "Too many arguments")
				return
			}
		}
	}
}

func main() {
	fmt.Println(gk.WithGreen().Bold("GOB:") + " go cross-platform builder tool")
	if len(os.Args) < 2 {
		fmt.Println("Usage: gob --option=")

		return
	}
	ParseArgs()

	if _, err := os.Open(inputFile); err != nil && inputFile != "." {
		fmt.Println(gk.Red("Error: ") + err.Error())
		os.Exit(1)
	}

	fmt.Println(gk.Green("Target OS: ") + targetOs)
	fmt.Println(gk.Green("Target Arch: ") + targetArch)
	fmt.Println(gk.Green("Output file: ") + out)
	fmt.Println(gk.Green("Input File: ") + inputFile)

}

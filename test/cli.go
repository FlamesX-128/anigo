package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FlamesX-128/anigo/src/types"
)

var Plugin = types.Plugin[types.Process]{
	Author: "FlamesX-128",
	Name:   "cli",
	This: types.Process{
		Handler: func(anigo *types.Anigo) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			fmt.Println(text)
		},
		Persistent: true,
	},
	Url:     "github.com/FlamesX-128/anigo-plugins/cli",
	Version: "0.1.0",
}

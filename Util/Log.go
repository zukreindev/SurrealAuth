package util

import (
	"fmt"
	"github.com/fatih/color"
)

func Log(title string, message string) {
	fmt.Printf("[" + color.GreenString(title) + "] " + message + "\n")
}
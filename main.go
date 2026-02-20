package main

import (
	"fmt"

	"github.com/Heribio/ValTracker/internal/jsonthings"
	"github.com/Heribio/ValTracker/internal/tui"
	"github.com/Heribio/ValTracker/internal/valorantapi"
)

func main() {
	if valorantapi.CheckToken() {
		tui.Run()
	} else {
		jsonthings.PromptToken()
		if valorantapi.CheckToken() {
			fmt.Println("API key working")
		} else {
			fmt.Println("API key not working")
		}
	}
}

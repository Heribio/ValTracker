package main

import (
	"github.com/Heribio/ValTracker/cmd"
	"github.com/Heribio/ValTracker/internal/tui"
)

func main() {
    cmd.Execute()
     tui.Run()
}

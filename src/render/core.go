package render

import (
	"github.com/nsf/termbox-go"
)

func Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Main event loop
	for {

	}
}

package pkg

import (
	"github.com/nsf/termbox-go"
)

func StartEditor() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
}

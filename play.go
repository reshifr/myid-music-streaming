package main

import (
	"fmt"
	"github.com/reshifr/play/core/codec"
)

func main() {
	music, ok := codec.GetFlac("build/sia.flac")
	if !ok {
		fmt.Println("ERROR!")
		return
	}
	fmt.Println(*music)
}

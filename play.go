package main

import (
	"fmt"
	"github.com/reshifr/play/core/codec"
)

func main() {
	music := codec.GetFlac("build/sia.flac")
	fmt.Println(*music)
}

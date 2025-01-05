package main

import (
	"fmt"
	"os"

	"github.com/ko6bxl/cm2img"
)

func main() {
	mode := os.Args[1]
	file := os.Args[2]
	fmt.Println(cm2img.Gen(mode, file))
}

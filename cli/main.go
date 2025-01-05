package main

import (
	"os"

	"github.com/ko6bxl/cm2img"
)

func main() {
	mode := os.Args[1]
	file := os.Args[2]
	cm2img.Gen(mode, file)
}

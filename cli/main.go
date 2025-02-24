package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ko6bxl/cm2img"
)

func main() {
	var mode string
	flag.StringVar(&mode, "m", "fine", "Mode to generate with: either 'fine' or 'normal'")
	var file string
	flag.StringVar(&file, "i", "USER IS BEING VERY DUMB", "File to convert into save, REQUIRED TO FUNCTION!")
	flag.Parse()

	if file == "USER IS BEING VERY DUMB" {
		fmt.Println("Input a file because file=" + file)
	} else {
		imgFile, err := os.Open(file)
		defer imgFile.Close()
		if err != nil {
			log.Fatal(err)
		}

		imageImage, err  := cm2img.GetImage(imgFile)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(cm2img.Gen(mode, imageImage))
	}

}

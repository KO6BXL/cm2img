package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"

	"github.com/nameless9000/cm2go/block"
	"github.com/nameless9000/cm2go/build"
)

func main() {
	var imgFile = os.Args[1]

	img, err := os.Open(imgFile)

	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()
	var supportedTypes = [...]string{
		".jpeg",
		".png",
		".jpg",
	}

	var isSupported = false

	for _, fileExt := range supportedTypes {
		if path.Ext(imgFile) == fileExt {
			isSupported = true
			break
		}
	}

	if !isSupported {
		log.Fatal(errors.New("ERROR: Unsupported file type. Please input a png or jpeg"))
	}

	Image, err := getImage(img)

	if err != nil {
		log.Fatal(err)
	}

	MaxX := Image.Bounds().Max.X
	MaxY := Image.Bounds().Max.Y

	var collection block.Collection
	var thingieidk *block.Base

	var bitshift uint32 = 1 << 4
	for ix := 0; ix < MaxX; ix++ {
		for iy := 0; iy < MaxY; iy++ {
			red, green, blue, alpha := Image.At(ix, iy).RGBA()

			if err != nil {
				log.Fatal(err)
			}

			var color2 block.Color

			color2.R = uint8(red / bitshift / bitshift)
			color2.G = uint8(green / bitshift / bitshift)
			color2.B = uint8(blue / bitshift / bitshift)
			a := uint8(alpha)

			if !(a < 15) {
				thingieidk = collection.Append(block.TILE(color2, 2))
				thingieidk.Offset.X = float32(ix)
				thingieidk.Offset.Z = float32(iy)
			}

		}
	}
	out, err := build.FastCompile([]block.Collection{collection})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
func getImage(img *os.File) (image.Image, error) {
	if path.Ext(img.Name()) == ".jpeg" || path.Ext(img.Name()) == ".jpg" {
		Image, err := jpeg.Decode(img)

		if err != nil {
			return nil, err
		}
		return Image, nil

	} else if path.Ext(img.Name()) == ".png" {
		Image, err := png.Decode(img)

		if err != nil {
			return nil, err
		}
		return Image, nil
	} else {
		return nil, errors.New("ERROR: Unsupported Image Type")
	}
}

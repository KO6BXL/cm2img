package cm2img

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"

	"github.com/nameless9000/cm2go/block"
	"github.com/nameless9000/cm2go/build"
)

func Gen(mode string, img image.Image) (string, error) {

	if mode == "normal" {
		out, err := build.FastCompile([]block.Collection{normMode(img)})

		if err != nil {
			return "", err
		}

		return out, nil
	} else if mode == "fine" {
		out, err := build.Compile([]block.Collection{fineMode(img)})

		if err != nil {
			return "", err
		}

		return out, nil
	}
	return "", errors.New("UNDEFINED ERROR")
}
func normMode(Image image.Image) block.Collection {

	MaxX := Image.Bounds().Max.X
	MaxY := Image.Bounds().Max.Y

	var collection block.Collection
	var thingieidk *block.Base

	var bitshift uint32 = 1 << 4
	for ix := 0; ix < MaxX; ix++ {
		for iy := 0; iy < MaxY; iy++ {
			red, green, blue, alpha := Image.At(ix, iy).RGBA()

			var color2 block.Color

			color2.R = uint8(red / bitshift / bitshift)
			color2.G = uint8(green / bitshift / bitshift)
			color2.B = uint8(blue / bitshift / bitshift)
			a := uint8(alpha)

			if !(a < 45) {
				thingieidk = collection.Append(block.TILE(color2, 2))
				thingieidk.Offset.X = float32(ix)
				thingieidk.Offset.Z = float32(iy)
			}

		}
	}
	return collection
}

func fineMode(Image image.Image) block.Collection {

	MaxX := Image.Bounds().Max.X
	MaxY := Image.Bounds().Max.Y

	var collection block.Collection
	var thingieidk *block.Base
	var yCount float32 = 0
	var yyCount int = 0
	var bitshift uint32 = 1 << 4
	for ix := 0; ix < MaxX; ix++ {
		yCount = yCount + 0.0005
		var z float32 = yCount

		for iy := 0; iy < MaxY; iy++ {
			red, green, blue, alpha := Image.At(ix, iy).RGBA()

			var color2 block.Color

			color2.R = uint8(red / bitshift / bitshift)
			color2.G = uint8(green / bitshift / bitshift)
			color2.B = uint8(blue / bitshift / bitshift)
			a := uint8(alpha / bitshift / bitshift)
			z = z - 0.0005

			if !(a < 45) {
				thingieidk = collection.Append(block.TILE(color2, 2))
				thingieidk.Offset.X = float32(float32(ix) / 5)
				thingieidk.Offset.Y = float32(z)
				thingieidk.Offset.Z = float32(float32(iy) / 5)
			}

		}
		yyCount = yyCount + 1

	}
	return collection
}

func GetImage(img *os.File) (image.Image, error) {
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

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"

	"git.neveris.one/gryffyn/snowcrash/utils"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 3 {
		log.Fatal("Usage: snowcrash <png file> <int x> <int y>")
	}
	x, _ := strconv.Atoi(arg[1])
	y, _ := strconv.Atoi(arg[2])
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	img := new(utils.Image)
	img.Open(arg[0])
	err := img.GetPixels()

	fmt.Print("PixelRGBA at 0,0 is: ")
	fmt.Print(img.PixelsRGBA[x][y])
	fmt.Println("\nHex: #" + utils.RgbaToHex(img.PixelsRGBA[x][y], false))

	if err != nil {
		log.Fatal("Error: Image could not be decoded")
	}

	fmt.Println(img.PixelsRGBA)

}

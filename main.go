package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strconv"

	"git.neveris.one/gryffyn/snowcrash/models/pixelSort"
	"git.neveris.one/gryffyn/snowcrash/utils"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 4 {
		log.Fatal("Usage: snowcrash <png file> <int x> <int y>")
	}
	_, _ = strconv.Atoi(arg[1])
	_, _ = strconv.Atoi(arg[2])
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	img := new(utils.Image)
	img.Open(arg[0])
	err := img.GetPixelsRGB()
	img.ToHSV()

	fmt.Println("Image dimensions: ")
	spew.Dump(img.Bounds)

	pixelSort.SortRows(img, pixelSort.Value)
	img.Write(arg[3])

	/* fmt.Print("PixelRGBA at " + strconv.Itoa(x) + "," + strconv.Itoa(y) + " is: ")
	fmt.Print(img.PixelsRGBA[x][y])
	fmt.Println("\nHex: #" + utils.RgbaToHex(img.PixelsRGBA[x][y], false))
	fmt.Print("PixelHSV at " + strconv.Itoa(x) + "," + strconv.Itoa(y) + " is: ")
	fmt.Print(img.Pixels.HSV[y][x]) */

	if err != nil {
		log.Fatal("Error: Image could not be decoded")
	}
}

package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"strconv"
)

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func main() {
	arg := os.Args[1:]
	x, _ := strconv.Atoi(arg[0])
	y, _ := strconv.Atoi(arg[1])
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open("./test/rgbg.png")

	if err != nil {
		log.Fatalln("Error: File could not be opened")
	}

	defer file.Close()

	pixels, err := getPixels(file)

	fmt.Print("Pixel at 0,0 is: ")
	fmt.Print(pixels[x][y])

	if err != nil {
		log.Fatal("Error: Image could not be decoded")
	}

	// fmt.Println(pixels)

}

func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)
	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)

	var pixels [][]Pixel
	var r, g, b, a uint8
	for y := 0; y < rect.Max.Y; y++ {
		var row []Pixel
		for x := 0; x < rect.Max.X; x++ {
			r = rgba.Pix[(y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4]
			g = rgba.Pix[((y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4)+1]
			b = rgba.Pix[((y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4)+2]
			a = rgba.Pix[((y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4)+3]
			row = append(row, rgbaToPixel(r, g, b, a))
		}
		pixels = append(pixels, row)
	}

	return pixels, err
}

func rgbaToPixel(r uint8, g uint8, b uint8, a uint8) Pixel {
	return Pixel{r, g, b, a}
}

func rgbaToHex(Pixel) string {

}

package utils

import (
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"strconv"
)

type PixelRGBA struct {
	R uint32
	G uint32
	B uint32
	A uint32
}

func (p PixelRGBA) RGBA() (r, g, b, a uint32) {
	return p.R, p.G, p.B, p.A
}

func NewRBGA(r, g, b, a uint32) PixelRGBA {
	return PixelRGBA{R: r, G: g, B: b, A: a}
}

type Image struct {
	File   io.Reader
	Image  image.Image
	Bounds image.Rectangle
	Pixels [][]PixelRGBA
}

func (i *Image) GetPixels() error {
	return i.GetPixelsRGB()
}

func (i *Image) GetPixelsRGB() error {
	img, _, err := image.Decode(i.File)
	i.Image = img
	i.Bounds = img.Bounds()

	var pixels [][]PixelRGBA
	var r, g, b, a uint32
	for y := i.Bounds.Min.Y; y < i.Bounds.Max.Y; y++ {
		var row []PixelRGBA
		for x := i.Bounds.Min.X; x < i.Bounds.Max.X; x++ {
			r, g, b, a = img.At(x, y).RGBA()
			row = append(row, rgbaToPixel(r, g, b, a))
		}
		pixels = append(pixels, row)
		/*
			for i := len(pixels)/2 - 1; i >= 0; i-- {
				opp := len(pixels) - 1 - i
				pixels[i], pixels[opp] = pixels[opp], pixels[i]
			}*/
	}
	i.Pixels = pixels
	return err
}

func (i *Image) Open(path string) {
	log.Println("Opening file " + path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Error: File could not be opened")
	}
	i.File = file
}

func (i *Image) Write(path string, img image.Image) {
	file, err := os.Create(path)
	err = png.Encode(file, img)
	if err != nil {
		log.Fatalln("file could not be written")
	}
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) PixelRGBA { return PixelRGBA{r, g, b, a} }

func Mph(hex string) string {
	if len(hex) == 1 {
		return "0" + hex
	}
	return hex
}

func RgbaToHex(p PixelRGBA, alpha bool) string {
	r := strconv.FormatInt(int64(p.R), 16)
	g := strconv.FormatInt(int64(p.G), 16)
	b := strconv.FormatInt(int64(p.B), 16)
	a := strconv.FormatInt(int64(p.A), 16)

	hex := Mph(r) + Mph(g) + Mph(b)
	if !alpha {
		return hex
	} else {
		return hex + Mph(a)
	}
}

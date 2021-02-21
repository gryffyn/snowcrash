package utils

import (
	"image"
	"image/draw"
	"io"
	"log"
	"os"
	"strconv"

	r2h "github.com/AlessandroPomponio/hsv/conversion"
)

type PixelRGBA struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type PixelHSV struct {
	H float64
	S float64
	V float64
}

type Image struct {
	File       io.Reader
	PixelsRGBA [][]PixelRGBA
	PixelsHSV  [][]PixelHSV
	Height     int
	Width      int
}

func (i *Image) GetPixels() error {
	return i.GetPixelsRGB()
}

func (i *Image) GetPixelsRGB() error {
	img, _, err := image.Decode(i.File)
	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)

	var pixels [][]PixelRGBA
	var r, g, b, a uint8
	for y := 0; y < rect.Max.Y; y++ {
		var row []PixelRGBA
		for x := 0; x < rect.Max.X; x++ {
			r = rgba.Pix[(y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4]
			g = rgba.Pix[((y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4)+1]
			b = rgba.Pix[((y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4)+2]
			a = rgba.Pix[((y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4)+3]
			row = append(row, rgbaToPixel(r, g, b, a))
		}
		pixels = append(pixels, row)
	}
	i.PixelsRGBA = pixels
	return err
}

func (i *Image) GetPixelsHSV() error {
	img, _, err := image.Decode(i.File)
	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rect, img, rect.Min, draw.Src)

	var pixels [][]PixelHSV
	var r, g, b, a uint8
	for y := 0; y < rect.Max.Y; y++ {
		var row []PixelHSV
		for x := 0; x < rect.Max.X; x++ {
			r = rgba.Pix[(y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4]
			g = rgba.Pix[((y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4)+1]
			b = rgba.Pix[((y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4)+2]
			a = rgba.Pix[((y-rect.Min.Y)*rgba.Stride+(x-rect.Min.X)*4)+3]
			row = append(row, RGBAToHSV(rgbaToPixel(r, g, b, a)))
		}
		pixels = append(pixels, row)
	}
	i.PixelsHSV = pixels
	return err
}

func (i *Image) ToHSV() {
	var pixels [][]PixelHSV
	for _, v := range i.PixelsRGBA {
		var row []PixelHSV
		for _, x := range v {
			row = append(row, RGBAToHSV(x))
		}
	}
	i.PixelsHSV = pixels
}

func (i *Image) Open(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Error: File could not be opened")
	}
	i.File = file
}

func rgbaToPixel(r uint8, g uint8, b uint8, a uint8) PixelRGBA { return PixelRGBA{r, g, b, a} }

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

func RGBAToHSV(p PixelRGBA) PixelHSV {
	var hsv PixelHSV
	hsv.H, hsv.S, hsv.V = r2h.RGBAToHSV(uint32(p.R), uint32(p.G), uint32(p.B), uint32(p.A))
	return hsv
}

package utils

import (
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"strconv"

	r2h "github.com/AlessandroPomponio/hsv/conversion"
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

type PixelHSV struct {
	H float64
	S float64
	V float64
}

type Pixels struct {
	RGBA [][]PixelRGBA
	HSV  [][]PixelHSV
}

type Image struct {
	File   io.Reader
	Image  image.Image
	Bounds image.Rectangle
	Pixels Pixels
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
		for i := len(pixels)/2 - 1; i >= 0; i-- {
			opp := len(pixels) - 1 - i
			pixels[i], pixels[opp] = pixels[opp], pixels[i]
		}
	}
	i.Pixels.RGBA = pixels
	return err
}

func (i *Image) GetPixelsHSV() error {
	err := i.GetPixelsRGB()
	i.ToHSV()
	return err
}

func (i *Image) ToHSV() {
	var pixels [][]PixelHSV
	for _, v := range i.Pixels.RGBA {
		var row []PixelHSV
		for _, x := range v {
			row = append(row, RGBAToHSV(x))
		}
		pixels = append(pixels, row)
	}
	i.Pixels.HSV = pixels
}

func (i *Image) Open(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Error: File could not be opened")
	}
	i.File = file
}

func (i *Image) Write(path string) {
	file, err := os.Create(path)

	img := overwritePixels(i)

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalln("file could not be written")
	}
}

func overwritePixels(i *Image) image.Image {
	nir := image.NewRGBA(i.Bounds)
	for y := i.Bounds.Min.Y; y < i.Bounds.Max.Y; y++ {
		for x := i.Bounds.Min.X; x < i.Bounds.Max.X; x++ {
			nir.Set(x, y, i.Pixels.RGBA[y][x])
		}
	}
	return nir
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

func RGBAToHSV(p PixelRGBA) PixelHSV {
	var hsv PixelHSV
	hsv.H, hsv.S, hsv.V = r2h.RGBAToHSV(p.R, p.G, p.B, p.A)
	return hsv
}

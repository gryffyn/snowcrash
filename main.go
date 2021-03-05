package main

import (
	"image"
	"image/png"
	"log"

	"git.neveris.one/gryffyn/snowcrash/models/glitch"
	"git.neveris.one/gryffyn/snowcrash/models/pixelSort"
	"git.neveris.one/gryffyn/snowcrash/utils"
)

func main() {

	pngName := "test/image0"

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	img := new(utils.Image)
	img.Open(pngName + ".png")
	err := img.GetPixelsRGB()

	// sort pixels
	_ = pixelSort.NewSorted(img)

	// diag stuff
	/* fmt.Print("PixelRGBA at " + strconv.Itoa(x) + "," + strconv.Itoa(y) + " is: ")
	fmt.Print(img.PixelsRGBA[x][y])
	fmt.Println("\nHex: #" + utils.RgbaToHex(img.PixelsRGBA[x][y], false))
	fmt.Print("PixelHSV at " + strconv.Itoa(x) + "," + strconv.Itoa(y) + " is: ")
	fmt.Print(img.Pixels.HSV[y][x]) */

	// test shifting pixels
	s := glitch.RandScaledShift(img.Bounds)
	log.Println(s)
	shifted := glitch.ShiftChannel(img.Image, 0, 3, 72)
	// shifted2 := glitch.ShiftChannel(shifted, 3, 1, s.X)

	// write final image
	img.Write(pngName+"_shifted.png", shifted) // shifted
	// img.Write(pngName+ "_out_sorted.png", sorted) // sorted

	log.Println("Written file")

	if err != nil {
		log.Fatal("Error: Image could not be decoded")
	}
}

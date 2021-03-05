package glitch

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	_ "git.neveris.one/gryffyn/snowcrash/utils"
)

type Shift struct {
	X int
	Y int
}

type RGBA struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func colToRGBA(c color.Color) RGBA {
	r, g, b, a := c.RGBA()
	return RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

// Shifts image direction by amount.
// Up:0 Down:1 Left:2 Right:3 / R:0 G:1 B:2
func ShiftChannel(i image.Image, chann uint8, direction uint8, shift int) image.Image {
	b := i.Bounds()
	nir := image.NewRGBA(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pixcol := i.At(x, y)
			col := colToRGBA(pixcol)
			if chann == 0 {
				if direction == 0 {
					col.R = colToRGBA(i.At(x, y-shift).(color.Color)).R
				} else if direction == 1 {
					col.R = colToRGBA(i.At(x, y+shift).(color.Color)).R
				} else if direction == 2 {
					col.R = colToRGBA(i.At(x-shift, y).(color.Color)).R
				} else if direction == 3 {
					col.R = colToRGBA(i.At(x+shift, y).(color.Color)).R
				}
			} else if chann == 1 {
				if direction == 0 {
					col.G = colToRGBA(i.At(x, y-shift).(color.Color)).G
				} else if direction == 1 {
					col.G = colToRGBA(i.At(x, y+shift).(color.Color)).G
				} else if direction == 2 {
					col.G = colToRGBA(i.At(x-shift, y).(color.Color)).G
				} else if direction == 3 {
					col.G = colToRGBA(i.At(x+shift, y).(color.Color)).G
				}
			} else if chann == 2 {
				if direction == 0 {
					col.B = colToRGBA(i.At(x, y-shift).(color.Color)).B
				} else if direction == 1 {
					col.B = colToRGBA(i.At(x, y+shift).(color.Color)).B
				} else if direction == 2 {
					col.B = colToRGBA(i.At(x-shift, y).(color.Color)).B
				} else if direction == 3 {
					col.B = colToRGBA(i.At(x+shift, y).(color.Color)).B
				}
			}
			nir.SetRGBA(x, y, color.RGBA(col))
		}
	}
	return nir
}

func RandScaledShift(r image.Rectangle) Shift {
	rand.Seed(time.Now().UnixNano())
	return Shift{
		X: rand.Intn((r.Max.X/20)-1) + 1,
		Y: rand.Intn((r.Max.Y/20)-1) + 1,
	}
}

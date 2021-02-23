package pixelSort

import (
	"math/rand"

	"git.neveris.one/gryffyn/snowcrash/utils"
)

type HSV int

const (
	Hue        HSV = iota
	Saturation HSV = iota
	Value      HSV = iota
)

/*
def threshold(image, lower_threshold, upper_threshold, **kwargs):
    intervals = []
    image_data = image.load()
    for y in range(image.size[1]):
        intervals.append([])
        for x in range(image.size[0]):
            level = lightness(image_data[x, y])
            if level < lower_threshold or level > upper_threshold:
                intervals[y].append(x)
    return intervals
*/

// Function is from above python code.
func threshold(i utils.Image, lower, upper int) [][]int {
	var intervals [][]int
	for y := 0; y < i.Bounds.Max.Y; y++ {
		intervals = append(intervals, []int{})
		for x := 0; x < i.Bounds.Max.X; y++ {
			val := i.Pixels.HSV[x][y].V
			if val < float64(lower) || val > float64(upper) {
				intervals[y] = append(intervals[y], x)
			}
		}
	}
	return intervals
}

// Quick sorts pixels based on given pixel value.
// 0 = Hue, 1 = Saturation, 2 = value
func QuickSort(a []utils.PixelHSV, sortType HSV) []utils.PixelHSV {
	if len(a) < 2 {
		return a
	}
	l, r := 0, len(a)-1
	p := rand.Int() % len(a)
	a[p], a[r] = a[r], a[p]

	for i := range a {
		if sortType == 0 {
			if a[i].H < a[r].H {
				a[l], a[i] = a[i], a[l]
				l++
			}
		}
		if sortType == 1 {
			if a[i].S < a[r].S {
				a[l], a[i] = a[i], a[l]
				l++
			}
		}
		if sortType == 2 {
			if a[i].V < a[r].V {
				a[l], a[i] = a[i], a[l]
				l++
			}
		}
	}

	a[l], a[r] = a[r], a[l]

	QuickSort(a[:l], sortType)
	QuickSort(a[l+1:], sortType)

	return a
}

func SortRows(i *utils.Image, sortType HSV) {
	var sorted [][]utils.PixelHSV
	var sortedRow []utils.PixelHSV
	for y := 0; y < i.Bounds.Max.Y; y++ {
		sortedRow = QuickSort(i.Pixels.HSV[y], sortType)
		sorted = append(sorted, sortedRow)
	}
	i.Pixels.HSV = sorted
}

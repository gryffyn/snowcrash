package pixelSort

import (
	"fmt"
	"math/rand"
	"sort"

	"git.neveris.one/gryffyn/snowcrash/utils"
	"github.com/davecgh/go-spew/spew"
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

func QSort(arr []utils.PixelHSV) []utils.PixelHSV {
	newArr := make([]utils.PixelHSV, len(arr))

	for i, v := range arr {
		newArr[i] = v
	}

	qsort(newArr, 0, len(arr)-1)

	return newArr
}

func qsort(arr []utils.PixelHSV, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i].V < pivot.V {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	qsort(arr, start, splitIndex-1)
	qsort(arr, splitIndex+1, end)
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
		if sortType == Hue {
			if a[i].H < a[r].H {
				a[l], a[i] = a[i], a[l]
				l++
			}
		}
		if sortType == Saturation {
			if a[i].S < a[r].S {
				a[l], a[i] = a[i], a[l]
				l++
			}
		}
		if sortType == Value {
			if a[i].V > a[r].V {
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

func SortRowsQuick(i *utils.Image) [][]utils.PixelRGBA {
	//spew.Dump(i.Pixels.HSV)
	var sorted [][]utils.PixelRGBA
	var sortedRow []utils.PixelRGBA
	for y := 0; y < i.Bounds.Max.Y; y++ {
		spew.Dump(i.Pixels.HSV[y])
		fmt.Print("---------------------------------------------------\n")
		sortedRow = utils.HSVArraytoRGBAArray(QSort(i.Pixels.HSV[y]))
		spew.Dump(sortedRow)
		sorted = append(sorted, sortedRow)
	}
	return sorted
}

func SortRows(i *utils.Image, sortType HSV) {
	var sorted [][]utils.PixelHSV
	grid := i.Pixels.HSV
	for y := 0; y < i.Bounds.Max.Y; y++ {
		for x := range grid[y] {
			sort.Slice(grid[y], func(i, j int) bool {
				switch sortType {
				case Hue:
					return grid[y][x].H < grid[y][x].H
				case Saturation:
					return grid[y][x].S < grid[y][x].S
				case Value:
					return grid[y][x].V < grid[y][x].V
				}
				return false
			})
			sorted = append(sorted, grid[y])
		}
	}
	i.Pixels.HSV = sorted
}

package pixelSort

import (
	"image"
	"sort"

	"git.neveris.one/gryffyn/snowcrash/utils"
)

// Returns brightness (sum of RGBA values)
func b(p utils.PixelRGBA) uint32 {
	return p.A + p.B + p.G + p.R
}

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
func threshold(i utils.Image, lower, upper uint32) [][]int {
	var intervals [][]int
	for y := 0; y < i.Bounds.Max.Y; y++ {
		intervals = append(intervals, []int{})
		for x := 0; x < i.Bounds.Max.X; y++ {
			val := b(i.Pixels[x][y])
			if val < lower || val > upper {
				intervals[y] = append(intervals[y], x)
			}
		}
	}
	return intervals
}

func QSort(arr []utils.PixelRGBA) []utils.PixelRGBA {
	newArr := make([]utils.PixelRGBA, len(arr))

	for i, v := range arr {
		newArr[i] = v
	}

	qsort(newArr, 0, len(arr)-1)

	return newArr
}

func qsort(arr []utils.PixelRGBA, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if b(arr[i]) < b(pivot) {
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

func SortRowsQuick(i *utils.Image) [][]utils.PixelRGBA {
	//spew.Dump(i.Pixels.HSV)
	var sorted [][]utils.PixelRGBA
	var sortedRow []utils.PixelRGBA
	for y := 0; y < i.Bounds.Max.Y; y++ {
		sortedRow = QSort(i.Pixels[y])
		sorted = append(sorted, sortedRow)
	}
	return sorted
}

func NewSorted(i *utils.Image) image.Image {
	sorted := SortRowsQuick(i)
	nir := image.NewRGBA(i.Bounds)
	for y := i.Bounds.Min.Y; y < i.Bounds.Max.Y; y++ {
		for x := i.Bounds.Min.X; x < i.Bounds.Max.X; x++ {
			nir.Set(x, y, sorted[y][x])
		}
	}
	return nir
}

func SortRows(i *utils.Image) {
	var sorted [][]utils.PixelRGBA
	grid := i.Pixels
	for y := 0; y < i.Bounds.Max.Y; y++ {
		for x := range grid[y] {
			sort.Slice(grid[y], func(i, j int) bool {
				return b(grid[y][x]) < b(grid[y][x])
			})
			sorted = append(sorted, grid[y])
		}
	}
	i.Pixels = sorted
}

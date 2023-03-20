package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)

	for y := range slice {
		slice[y] = make([]uint8, dx)
		for x := range slice[y] {
			slice[y][x] = uint8((x + y) / 2)
		}
	}
	return slice

}

func main() {
	pic.Show(Pic)
}

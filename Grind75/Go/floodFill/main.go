package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	tempColor := image[sr][sc]
	image[sr][sc] = color
	rDirections := []int{sr - 1, sr + 1}
	for _, direction := range rDirections {
		if direction < len(image) && direction >= 0 && image[direction][sc] == tempColor {
			image = floodFill(image, direction, sc, color)
		}
	}
	cDirections := []int{sc - 1, sc + 1}
	for _, direction := range cDirections {
		if direction < len(image[sr]) && direction >= 0 && image[sr][direction] == tempColor {
			image = floodFill(image, sr, direction, color)
		}
	}
	return image
}

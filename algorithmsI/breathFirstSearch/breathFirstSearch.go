package breathFirstSearch

// An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.
// You are also given three integers sr, sc, and newColor. You should perform a flood fill on the image starting from the pixel image[sr][sc].
// To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel,
// plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with newColor.
// Return the modified image after performing the flood fill.

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	color := image[sr][sc]

	if color == newColor {
		return image
	}
	dfs(image, sr, sc, color, newColor)
	return image
}

func dfs(image [][]int, r int, c int, color int, newColor int) {

	if image[r][c] == color {
		image[r][c] = newColor
		if r > 0 {
			dfs(image, r-1, c, color, newColor)
		}
		if c > 0 {
			dfs(image, r, c-1, color, newColor)
		}
		if r < len(image)-1 {
			dfs(image, r+1, c, color, newColor)
		}
		if c < len(image[0])-1 {
			dfs(image, r, c+1, color, newColor)
		}
	}

}

package breathFirstSearch

// An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.
// You are also given three integers sr, sc, and newColor. You should perform a flood fill on the image starting from the pixel image[sr][sc].
// To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel,
// plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with newColor.
// Return the modified image after performing the flood fill.
// Time Complexity: O(N), where N is the number of pixels in the image. We might process every pixel.
// Space Complexity: O(N), the size of the implicit call stack when calling dfs.
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
		if r-1 >= 0 {
			dfs(image, r-1, c, color, newColor)
		}
		if c-1 >= 0 {
			dfs(image, r, c-1, color, newColor)
		}
		if r+1 < len(image) {
			dfs(image, r+1, c, color, newColor)
		}
		if c+1 < len(image[0]) {
			dfs(image, r, c+1, color, newColor)
		}
	}

}

// You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
// The area of an island is the number of cells with a value 1 in the island.
// Return the maximum area of an island in grid. If there is no island, return 0.
// Time Complexity: O(R∗C), where RRR is the number of rows in the given grid, and C is the number of columns. We visit every square once.
// Space complexity: O(R∗C), the space used by seen to keep track of visited squares, and the space used by the call stack during our recursion.
func MaxAreaOfIsland(grid [][]int) int {
	result := 0

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 0 {
				continue
			}
			a := area(grid, r, c)
			if result < a {
				result = a
			}
		}
	}
	return result
}

func area(grid [][]int, r, c int) int {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0

	return 1 + area(grid, r, c-1) + area(grid, r, c+1) + area(grid, r+1, c) + area(grid, r-1, c)
}

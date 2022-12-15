package leetcode

// floodFill
// 733. Flood Fill
// https://leetcode.com/problems/flood-fill/
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	m := len(image) - 1
	n := len(image[0]) - 1

	cellColour := image[sr][sc]
	pixels := []int{sr, sc}

	for len(pixels) > 0 {
		x1, y1 := pixels[0], pixels[1]
		image[x1][y1] = color

		if x1 < m && image[x1+1][y1] == cellColour {
			pixels = append(pixels, x1+1, y1)
		}
		if x1 > 0 && image[x1-1][y1] == cellColour {
			pixels = append(pixels, x1-1, y1)
		}
		if y1 > 0 && image[x1][y1-1] == cellColour {
			pixels = append(pixels, x1, y1-1)
		}
		if y1 < n && image[x1][y1+1] == cellColour {
			pixels = append(pixels, x1, y1+1)
		}
		pixels = pixels[2:]

	}

	return image
}

// numIslands
// 200. Number of Islands
// https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	var islands int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				islands++
				destroyIsland(grid, i, j)
			}
		}
	}
	return islands
}

func destroyIsland(image [][]byte, sr int, sc int) {
	m := len(image) - 1
	n := len(image[0]) - 1

	cellColour := image[sr][sc]
	pixels := []int{sr, sc}

	for len(pixels) > 0 {
		x1, y1 := pixels[0], pixels[1]
		pixels = pixels[2:]

		if image[x1][y1] == '0' {
			continue
		}

		image[x1][y1] = '0'
		if x1 < m && image[x1+1][y1] == cellColour {
			pixels = append(pixels, x1+1, y1)
		}
		if x1 > 0 && image[x1-1][y1] == cellColour {
			pixels = append(pixels, x1-1, y1)
		}
		if y1 > 0 && image[x1][y1-1] == cellColour {
			pixels = append(pixels, x1, y1-1)
		}
		if y1 < n && image[x1][y1+1] == cellColour {
			pixels = append(pixels, x1, y1+1)
		}
	}
}

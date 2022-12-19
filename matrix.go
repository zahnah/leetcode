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

// maxAreaOfIsland
// 695. Max Area of Island
// https://leetcode.com/problems/max-area-of-island/
func maxAreaOfIsland(grid [][]int) int {
	var max, size int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				size = sizeOfIsland(grid, i, j)
				if size > max {
					max = size
				}
			}
		}
	}
	return max
}

func sizeOfIsland(image [][]int, sr int, sc int) int {
	size := 0
	m := len(image) - 1
	n := len(image[0]) - 1

	cellColour := image[sr][sc]
	pixels := []int{sr, sc}

	for len(pixels) > 0 {
		x1, y1 := pixels[0], pixels[1]
		pixels = pixels[2:]

		if image[x1][y1] == 0 {
			continue
		}

		image[x1][y1] = 0
		size++

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

	return size
}

// updateMatrix
// 542. 01 Matrix
// https://leetcode.com/problems/01-matrix/
func updateMatrix(mat [][]int) [][]int {
	result := make([][]int, len(mat))

	var q1, q2 []int

	for x := 0; x < len(mat); x++ {
		result[x] = make([]int, len(mat[x]))
		for y := 0; y < len(mat[x]); y++ {
			if mat[x][y] != 0 {
				q1 = append(q1, x, y)
				result[x][y] = -1
			}
		}
	}

	var min, step int
	for len(q1) > 0 {
		step++
		for len(q1) > 0 {
			x, y := q1[0], q1[1]
			q1 = q1[2:]
			min = -1
			if x > 0 && result[x-1][y] != -1 {
				if min == -1 || min > result[x-1][y] {
					min = result[x-1][y]
				}
			}
			if y > 0 && result[x][y-1] != -1 {
				if min == -1 || min > result[x][y-1] {
					min = result[x][y-1]
				}
			}
			if x < len(result)-1 && result[x+1][y] != -1 {
				if min == -1 || min > result[x+1][y] {
					min = result[x+1][y]
				}
			}
			if y < len(result[x])-1 && result[x][y+1] != -1 {
				if min == -1 || min > result[x][y+1] {
					min = result[x][y+1]
				}
			}
			if min == -1 || min > step {
				q2 = append(q2, x, y)
			} else {
				result[x][y] = min + 1
			}
		}
		q1 = q2
		q2 = []int{}
	}

	return result
}

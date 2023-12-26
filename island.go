package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])
	var numIslands int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				numIslands++

				// 从当前节点开始进行dfs遍历
				dfs(grid, i, j, m, n)
			}
		}
	}
	return numIslands
}

// 从二维数组的某一点开始进行dfs遍历
func dfs(grid [][]byte, i, j int, m, n int) {
	// 控制边界
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == '1' {
		// 将该节点置为0，表示被海水淹没
		grid[i][j] = '0'

		// 递归遍历上下左右四个方向
		dfs(grid, i-1, j, m, n)
		dfs(grid, i+1, j, m, n)
		dfs(grid, i, j-1, m, n)
		dfs(grid, i, j+1, m, n)
	}
	return
}

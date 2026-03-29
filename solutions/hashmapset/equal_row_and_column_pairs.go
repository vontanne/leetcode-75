package hashmapset

import (
	"strconv"
	"strings"
)

/*
2352. Equal Row and Column Pairs

Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj)
such that row ri and column cj are equal.

A row and column pair is considered equal if they contain the same elements in the
same order (i.e., an equal array).

Example 1:
	Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
	Output: 1
	Explanation: There is 1 equal row and column pair:
	- (Row 2, Column 1): [2,7,7]

Example 2:
	Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
	Output: 3
	Explanation: There are 3 equal row and column pairs:
	- (Row 0, Column 0): [3,1,2,2]
	- (Row 2, Column 2): [2,4,2,2]
	- (Row 3, Column 2): [2,4,2,2]

Constraints:
	- n == grid.length == grid[i].length
	- 1 <= n <= 200
	- 1 <= grid[i][j] <= 10^5
*/

func EqualPairs(grid [][]int) int {
	n := len(grid)
	rowCount := make(map[string]int, n)

	for i := range n {
		var sb strings.Builder
		for j, v := range grid[i] {
			if j > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(strconv.Itoa(v))
		}
		rowCount[sb.String()]++
	}

	count := 0

	for j := range n {
		var sb strings.Builder
		for i := range n {
			if i > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(strconv.Itoa(grid[i][j]))
		}
		count += rowCount[sb.String()]
	}

	return count
}

/*
Time Complexity: O(n^2)
	- Building row keys: O(n) rows * O(n) elements = O(n^2)
	- Building column keys and lookup: O(n) columns * O(n) elements = O(n^2)

Space Complexity: O(n^2)
	- rowCount map stores up to n keys, each key has length O(n)
	- strings.Builder internal buffer: O(n) per iteration, reused
*/

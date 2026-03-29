package hashmapset

/*
1207. Unique Number of Occurrences

Given an array of integers arr, return true if the number of occurrences of each value
in the array is unique or false otherwise.

Example 1:
	Input: arr = [1,2,2,1,1,3]
	Output: true
	Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1.
	No two values have the same number of occurrences.

Example 2:
	Input: arr = [1,2]
	Output: false

Example 3:
	Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
	Output: true

Constraints:
	- 1 <= arr.length <= 1000
	- -1000 <= arr[i] <= 1000
*/

func UniqueOccurrences(arr []int) bool {
	count := make(map[int]int)

	for _, num := range arr {
		count[num]++
	}

	seen := make(map[int]struct{})

	for _, occurrence := range count {
		if _, exists := seen[occurrence]; exists {
			return false
		}
		seen[occurrence] = struct{}{}
	}

	return true
}

/*
Time Complexity: O(n)
	- Counting occurrences: O(n)
	- Checking uniqueness: O(k) where k is number of unique elements, k <= n

Space Complexity: O(n)
	- count map stores up to n unique elements
	- seen map stores up to n unique occurrence counts
*/

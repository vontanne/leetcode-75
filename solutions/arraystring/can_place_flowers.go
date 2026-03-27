package arraystring

/*
605. Can Place Flowers

You have a long flowerbed in which some of the plots are planted, and some are not.
However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means
not empty, and an integer n, return true if n new flowers can be planted in the flowerbed
without violating the no-adjacent-flowers rule and false otherwise.

Example 1:
	Input: flowerbed = [1,0,0,0,1], n = 1
	Output: true

Example 2:
	Input: flowerbed = [1,0,0,0,1], n = 2
	Output: false

Constraints:
	- 1 <= flowerbed.length <= 2 * 10^4
	- flowerbed[i] is 0 or 1.
	- There are no two adjacent flowers in flowerbed.
	- 0 <= n <= flowerbed.length
*/

func CanPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	planted := 0

	for index := range length {
		if flowerbed[index] == 0 {
			leftEmpty := index == 0 || flowerbed[index-1] == 0
			rightEmpty := index == length-1 || flowerbed[index+1] == 0

			if leftEmpty && rightEmpty {
				flowerbed[index] = 1
				planted++
				if planted >= n {
					return true
				}
			}
		}
	}

	return planted >= n
}

/*
Time Complexity: O(n)
	- Single pass through the flowerbed array

Space Complexity: O(1)
	- Only using a few variables, modifying input array in place
*/

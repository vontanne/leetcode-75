package arraystring

/*
345. Reverse Vowels of a String

Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper
cases, more than once.

Example 1:
	Input: s = "IceCreAm"
	Output: "AceCreIm"
	Explanation: The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels,
	s becomes "AceCreIm".

Example 2:
	Input: s = "leetcode"
	Output: "leotcede"

Constraints:
	- 1 <= s.length <= 3 * 10^5
	- s consist of printable ASCII characters.
*/

func ReverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	chars := []byte(s)
	left := 0
	right := len(chars) - 1

	for left < right {
		for left < right && !vowels[chars[left]] {
			left++
		}
		for left < right && !vowels[chars[right]] {
			right--
		}

		if left < right {
			chars[left], chars[right] = chars[right], chars[left]
			left++
			right--
		}
	}

	return string(chars)
}

/*
Time Complexity: O(n)
	- Each character is visited at most once by either pointer

Space Complexity: O(n)
	- Converting string to byte slice (strings are immutable in Go)
	- Vowel map is O(1) - fixed size of 10 entries
*/

package main

import "fmt"

// Longest substring without repeating characters
func main() {
	len := lengthOfLongestSubstring("abbabcdaefghijklmnoprqvstuwzxy")
	fmt.Println(len)

}

func lengthOfLongestSubstring(s string) int {
	charLastIndex := make(map[rune]int)

	longestSubstringLength := 0  // ความยาว string ทั้งหมด
	currentSubstringLength := 0  // ความยาว string ปัจจุบัน
	start := 0

	for index, character := range s {
		// string ล่าสุด ต้องเท่ากับ string ใน[]
		lastIndex, ok := charLastIndex[character]
		// ถ้าไม่เท่า กรณีที่น้อยกว่าให้ เพิ่มไปเรื่อย ๆ
		if !ok || lastIndex < index-currentSubstringLength {
			currentSubstringLength++
		} else {
			//  ถ้ามากกว่า กำหนดให้เท่ากัน
			if currentSubstringLength > longestSubstringLength {
				longestSubstringLength = currentSubstringLength
			}
			// เริ่มที่ string ล่าสุด +1
			start = lastIndex + 1
			// ความยาว string ปัจจุบัน เท่ากับ ลำดับ ลบด้วย ค่าเริ่มต้น + 1
			currentSubstringLength = index - start + 1
		}
		// จะได้ string ที่ไม่ซ้ำเพิ่มลงใน []
		charLastIndex[character] = index
	}
	if currentSubstringLength > longestSubstringLength {
		longestSubstringLength = currentSubstringLength
	}
	return longestSubstringLength

}

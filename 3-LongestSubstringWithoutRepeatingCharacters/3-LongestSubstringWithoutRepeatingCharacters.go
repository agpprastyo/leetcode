// Last updated: 2/11/2026, 8:30:54 PM
func lengthOfLongestSubstring(s string) int {
    charIndexMap := make(map[byte]int)

    maxLength := 0

    start := 0

    for end := 0; end < len(s); end++ {
        currentChar := s[end]
        if lastIndex, found := charIndexMap[currentChar]; found && lastIndex >= start {
            start = lastIndex + 1
        }

        charIndexMap[currentChar] = end

        currentLength := end - start + 1

        if currentLength > maxLength {
            maxLength = currentLength
        }
    }
    return maxLength
    
}
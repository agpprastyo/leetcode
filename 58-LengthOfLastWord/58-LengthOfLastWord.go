// Last updated: 2/11/2026, 8:30:30 PM
func lengthOfLastWord(s string) int {
    n := len(s)

    i := n - 1

    for i >= 0 && s[i] == ' ' {
        i--
    }

    length := 0

    for i >= 0 && s[i] != ' ' {
        length++
        i--
    }

    return length

}
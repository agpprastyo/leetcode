func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }

    haystackLen := len(haystack)
    needleLen := len(needle)

    for i := 0; i<= haystackLen-needleLen; i++ {
        if haystack[i:i+needleLen] == needle {
            return i
        }
    }
    return -1
}
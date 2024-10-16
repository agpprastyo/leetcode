func reverse(x int) int {
    const INT_MAX = 1<<31 - 1 
	const INT_MIN = -1 << 31   

    reversed := 0
    for x != 0 {
        pop := x % 10
        x /= 10

        if reversed > (INT_MAX-pop)/10 || reversed < (INT_MIN-pop)/10 {
			return 0
		}
		reversed = reversed*10 + pop
    }

    return reversed
}
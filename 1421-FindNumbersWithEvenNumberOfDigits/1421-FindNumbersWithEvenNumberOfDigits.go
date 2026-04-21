// Last updated: 4/21/2026, 12:51:47 PM
func findNumbers(nums []int) int {
    evenCount := 0

    for _, num := range nums {
        digitCount := 0

        if num >= 10 && num <= 99 {
            digitCount = 2
        } else if num >= 1000 && num <= 9999 {
            digitCount = 4
        } else if num == 100000 {
            digitCount = 6
        } else {
            digitCount = 1
        }

        if digitCount % 2 == 0 {
            evenCount++
        }
    }

    return evenCount
}
class Solution {
    fun romanToInt(s: String): Int {
         val romanMap = mapOf(
        'I' to 1,
        'V' to 5,
        'X' to 10,
        'L' to 50,
        'C' to 100,
        'D' to 500,
        'M' to 1000
    )

    var total = 0
    var prevValue = 0

    for (char in s) {
        val currentValue = romanMap[char] ?: 0

        if (currentValue > prevValue) {
            // Subtract the previous value twice (since it was added once before)
            total += currentValue - 2 * prevValue
        } else {
            total += currentValue
        }

        prevValue = currentValue
    }

    return total
    }
}
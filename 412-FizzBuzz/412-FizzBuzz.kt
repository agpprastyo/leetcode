// Last updated: 2/11/2026, 8:30:20 PM
class Solution {
    fun fizzBuzz(n: Int): List<String> {
        return(1..n).map { it ->
            when {
                it % 15 == 0 -> "FizzBuzz"
                it % 3 == 0 -> "Fizz"
                it % 5 == 0 -> "Buzz"
                else -> it.toString()
            }

        }
    }
}
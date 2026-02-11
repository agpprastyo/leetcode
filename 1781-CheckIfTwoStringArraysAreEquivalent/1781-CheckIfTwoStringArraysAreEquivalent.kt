// Last updated: 2/11/2026, 8:30:14 PM
class Solution {
    fun arrayStringsAreEqual(word1: Array<String>, word2: Array<String>): Boolean {
        val concatenatedWord1 = word1.joinToString("")
        val concatenatedWord2 = word2.joinToString("")

        return concatenatedWord1 == concatenatedWord2
    }
}
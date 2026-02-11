// Last updated: 2/11/2026, 8:30:11 PM
class Solution {
    fun mergeAlternately(word1: String, word2: String): String {
        val merged = StringBuilder()
        val maxLength = maxOf(word1.length, word2.length)

        for (i in 0 .. maxLength) {
            if (i < word1.length) {
                merged.append(word1[i])
            }
            if (i < word2.length) {
                merged.append(word2[i])
            }
        }
        return merged.toString()
    }
}
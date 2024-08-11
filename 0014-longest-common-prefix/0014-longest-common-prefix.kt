class Solution {
    fun longestCommonPrefix(strs: Array<String>): String {
         if (strs.isEmpty()) return ""

    var prefix = strs[0]  // Start with the first string as the prefix

    for (i in 1 until strs.size) {
        while (strs[i].indexOf(prefix) != 0) {
            // Reduce the prefix length until it matches the start of strs[i]
            prefix = prefix.substring(0, prefix.length - 1)
            if (prefix.isEmpty()) return ""
        }
    }

    return prefix
    }
}
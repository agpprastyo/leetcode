// Last updated: 2/11/2026, 8:30:15 PM
class Solution {
    fun runningSum(nums: IntArray): IntArray {
       var i = 1
        while(i < nums.count()) {
            nums[i] += nums[i-1]
            i++
        }
        return nums
    }
}
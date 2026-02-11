// Last updated: 2/11/2026, 8:30:37 PM
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    k := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[k-1] {
            nums[k] = nums[i]
            k++
        }
    }
    return k
}
// Last updated: 2/11/2026, 8:30:35 PM
func removeElement(nums []int, val int) int {
    k := 0


    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[k] = nums[i]
            k++ 
        }
    }

    return k
}
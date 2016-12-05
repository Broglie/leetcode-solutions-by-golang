// 解法一：效率较高，O(n)
 func majorityElement(nums []int) int {
    major := nums[0]
    count := 1
    for i := 1; i < len(nums); i++ {
        if count == 0 {
            major = nums[i]
            count++
        } else if major == nums[i] {
            count++
        } else {
            count--
        }
    }
    return major
}

// 解法二：效率较低
func majorityElement(nums []int) int {
    sort.Ints(nums)
    println(nums[len(nums)/2])
}
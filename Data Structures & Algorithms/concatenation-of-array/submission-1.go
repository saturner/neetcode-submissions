func getConcatenation(nums []int) []int {
    var ans []int
    ans = append(ans, nums...)
    ans = append(ans, nums...)
    return ans
}

func getMaximumXor(nums []int, maximumBit int) []int {
    xors := make([]int, len(nums))
    result := make([]int, len(nums))

    maxValue := int(math.Pow(2, float64(maximumBit))) - 1

    xors[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        xors[i] = nums[i] ^ xors[i-1]
    } 

    for i := 0; i < len(nums); i++ {
        result[i] = maxValue - xors[len(nums)-i-1]
    } 

    return result
}

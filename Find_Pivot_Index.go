func pivotIndex(nums []int) int {
    temp := make([]int,len(nums)+1)
    sums := Sum(nums)
    for indx, val := range nums{
        if indx == 0 {
            sums -= val
            temp[indx]=0
        }else{
            temp[indx] = temp[indx-1]+nums[indx-1]
            sums -= val 
        }
        if sums == temp[indx] {
            return indx
        }
    }
    return -1
}


func Sum(nums []int) int{
    var sum int
    for _,num := range nums {
        sum +=num
    }
    return sum
}

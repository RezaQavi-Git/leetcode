package problems

func binarySearch(nums []int, l, h, target int) int {
    if l > h {
        return l
    }

    mid := (l + h) / 2

    if nums[mid] == target {
        return mid
    } else if nums[mid] < target {
        return binarySearch(nums, mid+1, h, target)
    } else {
        return binarySearch(nums, l, mid-1, target)
    }
}

func SearchInsert(nums []int, target int) int {
    return binarySearch(nums, 0, len(nums)-1, target)
}

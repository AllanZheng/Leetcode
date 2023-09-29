package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	result := []int{}
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	if m < n {
		result = append(result, nums2[j:]...)
	} else {
		result = append(result, nums1[i:]...)
	}
	copy(nums1, result)
	return

}
func main() {

}

package Test_password

import (
	"fmt"
	"testing"
)

func Test_password(t *testing.T) {
	var nums1 []int
	nums2 := []int{1, 2, 3}
	nums3 := append(nums1, nums2...)
	fmt.Println(len(nums3))
}

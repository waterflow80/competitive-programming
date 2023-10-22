package main

import (
	"fmt"
	"time"
)

func main() {
	arr1 := []int{1,2,3, 4}
	arr2 := []int{4,5,6,7,8, 9}
	t1 := time.Now()
	median := findMedianSortedArrays(arr1, arr2)
	t2 := time.Now()
	fmt.Printf("The median of %v and %v is: %v\n", arr1, arr2, median)
	fmt.Printf("Total time of execution is %v\n", t2.Sub(t1))
}


/**
Return the median of two sorted arrays
(This was my intuitive approach, and it can be defintely
	enhanced)*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var temp []int
	
	n1 := len(nums1)
	n2 := len(nums2) 

	/// Making both arrays of the same size
	if n1 < n2 {
		for i:=0; i<(n2-n1);i++ {
			nums1 = append(nums1, 9223372036854775807)
		}
	}else {
		for i:=0; i<(n1-n2);i++ {
			nums2 = append(nums2, 9223372036854775807)
	}
}
	// And adding another value at the end
	nums1 = append(nums1, 9223372036854775807)
	nums2 = append(nums2, 9223372036854775807)
	
	n := n1 + n2 // Total number of elements
	if (n % 2 == 0) { 
			///// The total number of elements is even /////

			mx_indx := n / 2 // max index of the resulting array
			f := 0 // loop over nums1
			k := 0 // loop over nums2
			for ; (f+k) <= mx_indx; {
					if nums1[f] < nums2[k] {
							temp = append(temp, nums1[f])
							f++
					} else {
							temp = append(temp, nums2[k])
							k++
					}
			}
			return (float64(temp[mx_indx]) + float64(temp[mx_indx-1])) / 2
	} else {
			///// The total number of elements is odd /////
			mx_indx := n / 2 // floor value
			f := 0 // loop over nums1
			k := 0 // loop over nums2
			for ; (f+k) <= mx_indx; {
					if nums1[f] < nums2[k] {
							temp = append(temp, nums1[f])
							f++
					} else {
							temp = append(temp, nums2[k])
							k++
					}
			}
			return float64(temp[len(temp)-1])
	}
}
package main

type SparseVector struct {
	Vector []int
}

//Use hasmap = option 2
func Constructor(nums []int) SparseVector {

	// var arr SparseVector

	test := make([]int, len(nums))
	test = nums
	return SparseVector{
		Vector: test,
	}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {

	//product := make([]int,len(this.Vector))
	product := 0
	for i := 0; i < len(this.Vector); i++ {
		//fmt.Println("this.Vector[i]",this.Vector[i])
		// fmt.Println("vec.Vector[i]",vec.Vector[i])
		product += this.Vector[i] * vec.Vector[i]
	}
	// fmt.Println(product)
	return product
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */

package backtracking

// Given an array nums of distinct integers, return all the possible permutations.
// You can return the answer in any order.
// Time complexity : k-permutations_of_n, or partial permutation.
// Space complexity : O(N!) since one has to keep N! solutions.
func Permute(nums []int) [][]int {
	var result [][]int

	permuteBacktrack(len(nums), 0, nums, &result)

	return result
}

func permuteBacktrack(length int, index int, nums []int, result *[][]int) {

	if index == length {
		*result = append(*result, nums)
		return
	}

	for i := index; i < length; i++ {
		t := make([]int, length)
		copy(t, nums)
		t[index], t[i] = t[i], t[index]
		permuteBacktrack(length, index+1, t, result)
	}
}

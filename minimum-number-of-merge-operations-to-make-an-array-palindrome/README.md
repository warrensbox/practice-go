Find the minimum number of merge operations to make an array palindrome

Given a list of non-negative integers, find the minimum number of merge operations to make it a palindrome. A merge operation can only be performed on two adjacent elements. The result of a merge operation is that the two adjacent elements are replaced with their sum.

For example,

Input:  [6, 1, 3, 7]
 
Output: 1
 
Explanation: [6, 1, 3, 7] —> Merge 6 and 1 —> [7, 3, 7]
 
 
Input:  [6, 1, 4, 3, 1, 7]
 
Output: 2
 
Explanation: [6, 1, 4, 3, 1, 7] —> Merge 6 and 1 —> [7, 4, 3, 1, 7] —> Merge 3 and 1 —> [7, 4, 4, 7]
 
 
Input:  [1, 3, 3, 1]
 
Output: 0
 
Explanation: The list is already a palindrome
Practice this problem

We can easily solve the problem by maintaining two pointers, i and j, that initially points to both endpoints of array arr, i.e., the first pointer i points to the index of the first array element and the second pointer j points to the index of the last array element. Then loop till the search space is exhausted and reduce search space arr[i, j] at each iteration of the loop.

In each iteration of the loop, we compare the elements present at index i and j and perform the merge operation on the lesser weight side, i.e., merge element arr[i] with element arr[i+1] if arr[i] < arr[j] and increment i, merge element arr[j] with element arr[j-1] when arr[i] > arr[j] and decrement j; otherwise, ignore both the elements when arr[i] == arr[j].


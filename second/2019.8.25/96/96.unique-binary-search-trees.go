package problem96

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */
//  使用递归的思想，大问题分解
func numTrees(n int) int {
	//  返回条件
	if n == 0 {
		return 1
	}

	if n == 1 || n == 2 {
		return n
	}

	if n == 3 {
		return 5
	}

	res := 0
	//  递推公式，选取其中一个作为 root节点，子问题依旧是这样。
	//  考虑到对称结构，只算了一半
	for i := 1; i <= n/2; i++ {
		res += numTrees(i-1) * numTrees(n-i)
	}
	res *= 2

	// n为奇数是，中间节点在上述循环中没有考虑进去
	if n%2 == 1 {
		temp := numTrees(n / 2)
		res += temp * temp
	}
	return res

}

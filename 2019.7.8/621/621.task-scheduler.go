package problem621

/*
 * @lc app=leetcode id=621 lang=golang
 *
 * [621] Task Scheduler
 */
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	res := [26]int{}
	//  统计每个任务出现的次数
	for _, c := range tasks {
		res[c-'A']++
	}

	//  寻找其中出现次数最多的任务

	most := 0
	for _, num := range res {
		if num > most {
			most = num
		}
	}

	// 至少存在idle 的数量
	idles := (most - 1) * (n + 1)

	for _, num := range res {
		idles -= min(most-1, num)
	}
	return len(tasks) + max(0, idles)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

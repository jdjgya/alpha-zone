package main

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxDistToClosest(seats []int) int {
	n := len(seats)
	start := -1
	maxGap := 0

	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			if start == -1 {
				maxGap = i
			} else {
				maxGap = max(maxGap, (i-start)/2)
			}

			start = i
		}
	}

	maxGap = max(maxGap, n-1-start)

	return maxGap
}

func main() {
	maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1})
	maxDistToClosest([]int{1, 0, 0, 0})
	maxDistToClosest([]int{0, 0, 0, 1})
	maxDistToClosest([]int{0, 1})
}

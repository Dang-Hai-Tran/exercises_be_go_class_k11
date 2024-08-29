package main

func FindSliceMinMaxAvg(list []int) (int, int, float64) {
	if len(list) == 0 {
		return 0, 0, 0
	}
	min := list[0]
	max := list[0]
	sum := 0
	for _, v := range list {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}
	avg := float64(sum) / float64(len(list))
	return min, max, avg
}

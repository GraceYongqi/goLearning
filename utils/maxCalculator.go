package utils

// calculateMax 计算最大值
func CalculateMax(a []int) int {
	max := 0
	for _, value := range a {
		if value > max {
			max = value
		} else {
			//do nothing
		}
	}
	return max
}
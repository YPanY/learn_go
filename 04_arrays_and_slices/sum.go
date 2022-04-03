package arrays

func Sum(numbers [5]int) int{
	ans := 0
	for i :=0; i < 5 ; i++{
		ans+=numbers[i]
	}
	return ans
}
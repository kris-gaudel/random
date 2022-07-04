package random

func MiddleSquare(seed int) []int {
	randomNums := make([]int, 20)

	n := seed * seed / 1000
	
	for i := 0; i <= 19; i++ {
		n = n * n / 1000
		randomNums[i] = n
	}

	return randomNums
}

func LinearCongruentialGenerator(a int, x_n int, c int, m int) []int {
	randomNums := make([]int, 20)

	rand := (a * x_n + c) % m
	randomNums[0] = rand

	for i := 1; i <= 19; i++ {
		rand = (a * rand + c) % m
		randomNums[i] = rand
	}

	return randomNums
}	

func Xorshift(seed int) []int {
	randomNums := make([]int, 20)

	for i := 0; i <= 19; i++ {
		seed = seed << 13
		seed = seed >> 17
		seed = seed << 5
		randomNums[i] = seed
	}

	return randomNums
}
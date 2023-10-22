package main

func encontrarSegundoMaior(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	maior := nums[0]
	segundoMaior := nums[1]

	for i := 1; i < len(nums); i++ {
		if nums[i] > maior {
			segundoMaior = maior
			maior = nums[i]
		} else if nums[i] > segundoMaior {
			segundoMaior = nums[i]
		}
	}

	return segundoMaior
}

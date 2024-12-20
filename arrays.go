package main

func ArraySolutions(nums []int, n int, question int) any {
	switch question {
	case 1:
		//1431. Kids With the Greatest Number of Candies
		return kidsWithCandies(nums, n)
	case 2:
		// 605. Can Place Flowers
		return canPlaceFlowers(nums, n)
	case 3:
		// 238. Product id Array Except Self
		return productExceptSelf(nums)
	case 4:
		//283. Move Zeroes
		return moveZeroes(nums)
	case 5:
		//643. Maximum Average Subarray I
		return findMaxAverage(nums, n)

	}
	return "invalid"
}

//1431. Kids With the Greatest Number of Candies
func kidsWithCandies(candies []int, extraCandies int) []bool {
	var response []bool
	m := 0
	for i := 0; i < len(candies); i++ {
		m = max(m, candies[i])
	}
	for i := range candies {
		if candies[i]+extraCandies >= m {
			response = append(response, true)
		} else {
			response = append(response, false)
		}
	}
	return response
}

// 605. Can Place Flowers
func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	i := 0
	for i < length {
		if flowerbed[i] == 0 {
			prev := (i == 0) || (flowerbed[i-1] == 0)
			next := (i == length-1) || (flowerbed[i+1] == 0)
			if prev && next {
				n--
				flowerbed[i] = 1
			}
		}
		i++

	}
	return (n <= 0)

}

// 238. Product id Array Except Self
func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	output := make([]int, n)

	left := 1
	for i := 0; i < n; i++ {
		output[i] = left
		left *= nums[i]
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= right
		right *= nums[i]
	}

	return output

}

//283. Move Zeroes
func moveZeroes(nums []int) []int {
	var count int
	if len(nums) > 1 {
		for i := range nums {
			if nums[i] == 0 {
				count++
			}
		}
		for count > 0 {
			for i := range nums {
				if nums[i] == 0 && i < len(nums)-1 {
					temp := nums[i+1]
					nums[i] = temp
					nums[i+1] = 0
				}

			}
			count--
		}

	}
	return nums

}

//643. Maximum Average Subarray I
func findMaxAverage(nums []int, k int) float64 {
	sum := 0.0
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	maxSum := sum

	for i := k; i < len(nums); i++ {
		sum += float64(nums[i]) - float64(nums[i-k])
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum / float64(k)
}

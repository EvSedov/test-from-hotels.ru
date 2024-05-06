package main

import (
	"fmt"
	"math"
)

// Первая задача
func Declension(n int) string {
	switch n % 10 {
	case 1:
		if n%100 != 11 {
			return fmt.Sprintf("%d компьютер", n)
		}
	case 2, 3, 4:
		if n%100 != 12 && n%100 != 13 && n%100 != 14 {
			return fmt.Sprintf("%d компьютера", n)
		}
	}
	return fmt.Sprintf("%d компьютеров", n)
}

// Вторая задача
func CommonDivisors(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	commonDivisors := []int{}
	for i := 2; i <= min; i++ {
		isCommon := true

		for _, num := range nums {
			if num%i != 0 {
				isCommon = false
				break
			}
		}

		if isCommon {
			commonDivisors = append(commonDivisors, i)
		}
	}

	return commonDivisors
}

// Третья задача
func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findPrimes(min, max int) []int {
	p := []int{}

	for i := min; i <= max; i++ {
		if isPrime(i) {
			p = append(p, i)
		}
	}

	return p
}

// Четвертая задача
func multiplyTable(n int) {
	fmt.Printf("   ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("%2d ", i)
		for j := 1; j <= n; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	result1 := Declension(1)
	fmt.Println(result1)
	result1 = Declension(1244)
	fmt.Println(result1)
	result1 = Declension(25)
	fmt.Println(result1)
	fmt.Println()

	nums := []int{42, 12, 18}
	result2 := CommonDivisors(nums)
	fmt.Println(result2)
	fmt.Println()

	min := 11
	max := 20
	result3 := findPrimes(min, max)
	fmt.Println(result3)
	fmt.Println()

	multiplyTable(3)
	multiplyTable(7)
	multiplyTable(15)
}

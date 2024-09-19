package strongnumber

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func Strong(n int) string {

	digit := 0
	result := 0
	originalN := n

	for n > 0 {
		digit = n % 10
		result += fact(digit)
		n = n / 10
	}
	fmt.Println(result)

	if originalN == result {
		return "STRONG!!!!"
	} else {
		return "Not Strong !!"
	}

}

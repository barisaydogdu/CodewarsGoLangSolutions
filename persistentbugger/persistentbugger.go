package persistentbugger

func Persistence(n int) int {
	counter := 0
	//if n is lower than 10 return 0
	if n < 10 {
		return 0
	}
	//loop up to single digit number
	for n >= 10 {
		result := 1
		for n > 0 {
			//last digit
			digit := n % 10
			//multiply all digits
			result *= digit

			n = n / 10
		}
		//update n value
		n = result
		//count multiplication
		counter++
	}
	return counter
}

package pkg

import "strconv"

/* FizzBuzz
 Returns "Fizz" if the num is multiple of 3, "Buzz" if the num is multiple of 5,
 "FizzBuzz" if the num is multiple of 3 and 5 else return the num
* @param num int
 returns string
*/
func FizzBuzz(num int) string {
	if num == 0 {
		return "0"
	}
	if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	if num%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(num)
}

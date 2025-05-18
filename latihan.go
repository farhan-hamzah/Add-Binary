package main
import (
	"fmt"
)

func addBinary(a string, b string) string {
	result := ""
	carry := 0

	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			if a[i] == '1' {
				sum += 1
			}
			i--
		}

		if j >= 0 {
			if b[j] == '1' {
				sum += 1
			}
			j--
		}

		// Tambahkan digit hasil ke depan string result
		result = string('0'+(sum%2)) + result

		// Update carry
		carry = sum / 2
	}

	return result
}

func main() {
	fmt.Println(addBinary("11", "1"))       // Output: "100"
	fmt.Println(addBinary("1010", "1011"))  // Output: "10101"
	fmt.Println(addBinary("0", "0"))        // Output: "0"
	fmt.Println(addBinary("111", "111"))    // Output: "1110"
}

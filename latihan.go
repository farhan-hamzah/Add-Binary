// package main
// import "fmt"
// const NMAX int = 100
// func main(){
// 	var A[NMAX]int
// 	var n, i int
// 	fmt.Scan(&n)
// 	for i = 0; i < n; i++{
// 		fmt.Scan(&A[i])
// 	}
// 	hasil := avaregValue(A[:n], n)
// 	fmt.Print(hasil)
// }
// func avaregValue(nums [] int, n int)int{
// 	var i, jum, c int
// 	for i = 0; i < n; i++{
// 		if nums[i]%2 == 0 && nums[i]%3 == 0{
// 			jum+= nums[i]
// 			c++
// 		}
// 	}
// 	if c == 0{
// 		return 0
// 	}else{
// 		return jum/c
// 	}
// }


// package main

// import (
// 	"fmt"
// )

// func lengthOfLastWord(s string) int {
// 	n := len(s)
// 	length := 0
// 	i := n - 1

// 	// Lewati spasi di akhir string
// 	for i >= 0 && s[i] == ' ' {
// 		i--
// 	}

// 	// Hitung panjang kata terakhir
// 	for i >= 0 && s[i] != ' ' {
// 		length++
// 		i--
// 	}

// 	return length
// }

// func main() {
// 	var word string
// 	fmt.Scan(&word)
// 	fmt.Print(lengthOfLastWord(word))
// }



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

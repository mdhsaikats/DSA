package array

import "fmt"

func Insertion(n, pos int) {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		if i == pos {
			arr[pos] = n
			fmt.Println(arr[i])
		} else {
			fmt.Println(arr[i])
		}
	}
}

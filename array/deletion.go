package array

import "fmt"

func Deletion(pos int) {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		if i == pos {
			arr[pos] = 0
			fmt.Println(arr[i])
		} else {
			fmt.Println(arr[i])
		}
	}
}

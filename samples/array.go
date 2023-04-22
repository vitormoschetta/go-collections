package samples

import (
	"fmt"
	"sort"
)

func ArraySample() {
	var arr [10]string

	for i := 0; i < len(arr); i++ {
		key := i
		value := getRandonString()
		arr[key] = value
	}

	fmt.Println(arr)

	sort.Strings(arr[:])

	fmt.Println(arr)
}

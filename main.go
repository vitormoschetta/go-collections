package main

import (
	"fmt"
	"sort"

	"collections/samples"
)

func main() {
	// arraySample()
	// sliceSample()
	// mapSample()
	samples.ArraySample()
}

func arraySample() {
	// declare array
	var arr [3]string

	// add value to array
	for i := 0; i < len(arr); i++ {
		value := fmt.Sprintf("value%d", i)
		arr[i] = value
	}

	// update value
	arr[0] = "value100"

	// delete value
	arr[0] = "" // note: array is fixed size
	fmt.Println(arr)
}

func sliceSample() {
	// declare slice
	var slice []string

	// add value to slice
	for i := 0; i < 3; i++ {
		value := fmt.Sprintf("value%d", i)
		slice = append(slice, value)
	}

	// update value
	slice[0] = "value100"

	// delete value
	slice = append(slice[:0], slice[1:]...) // delete first element
	fmt.Println(slice)

	// sort slice
	sort.Strings(slice)
}

func mapSample() {
	// declare map
	var m map[int]string = make(map[int]string)

	// add value to map
	for i := 0; i < 3; i++ {
		key := i
		value := fmt.Sprintf("value%d", i)
		m[key] = value
	}

	// update value
	m[0] = "value100"

	// delete value
	delete(m, 0)
	fmt.Println(m)
}

package internal

import "fmt"

func Arrays() {
	var intArray [3]int32
	intArray[0] = 123

	fmt.Println(intArray[0])
	fmt.Println(&intArray[0])

	intArray2 := [3]int32{1, 2, 3}

	rangeFromIndex1ToIndex2 := intArray2[1:3]

	fmt.Println(rangeFromIndex1ToIndex2)

	var intSlice []int32 = []int32{1, 2, 3}
	intSlice = append(intSlice, 23)

	//type length and capacity
	var intSlice2 []int32 = make([]int32, 4, 8)
	intSlice2 = append(intSlice2, intSlice...)

	fmt.Println(intSlice2)

	var map1 map[string]uint8 = make(map[string]uint8)
	fmt.Println(map1)

	var map2 = map[string]uint8{"key1": 23, "key2": 32}

	fmt.Println(map2["key1"])
	fmt.Println(map2["key2"])

	var key1Vlaue, ok = map2["key1"]

	if ok {
		fmt.Printf("key1 exists in map. value: %v\n", key1Vlaue)
	}

	fmt.Println("iterating over a map... ")
	for key, value := range map2 {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	fmt.Println("iterating over an array...")
	for index, value := range intSlice2 {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
}

package internal

import "fmt"

type structTest struct {
	v1 uint16
	v2 uint16
}

func (st structTest) multiplyProperties() uint16 {
	return st.v1 + st.v2
}

func Structs() {
	st1 := structTest{v1: 25, v2: 25}

	fmt.Println(st1.multiplyProperties())
}

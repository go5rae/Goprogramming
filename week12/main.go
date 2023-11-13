package main

import "fmt"

// import "fmt"

func main() {
	a := make([]string, 4, 5) // type, length, capacity(수용량)
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	as[1] = "Z"
	c := append(a, "y")
	// c := append(a, "y", "x")

	c[0] = "q"
	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))
	fmt.Printf("%x %x %x\n", &a[0], &as[0], &c[0])
}

// func main() {
// 	a := []string{"a", "b", "c", "d"}
// 	as := a[:2]
// 	as[1] = "Z"
// 	fmt.Println(a, as)

// 	b := [4]int{4, 3, 2, 1}
// 	bs := b[1:3]
// 	fmt.Println(b, bs)
// }

// package main

// import "fmt"

// func main() {
// 	s := []int{0, 0, 0, 0, 0} // 단축연산자 사용, 변수 선언과 동시에 메모리 할당 그리고 원소 초기화

// 	s[4] = 100
// 	s[0] = 7
// 	s[1] = 91

// 	for _, value := range s {
// 		fmt.Println(value)
// 	}

// 	copyS := s[1:4]
// 	for _, value := range copyS {
// 		fmt.Println(value)
// 	}

// 	test := [3]string{"inha", "go", "student"}
// 	testS := test[:2] // testS := test[0:2]
// 	testS2 := test[1:]

// 	testS2[0] = "python"
// 	// test[1] = "python"
// 	// testS[1] = "python"

// 	fmt.Println(test, len(test))
// 	fmt.Println(testS, len(testS))
// 	fmt.Println(testS2, len(testS2))
// }

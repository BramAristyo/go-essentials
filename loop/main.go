package main

import "fmt"

func main()  {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}	

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	nums := []int{10, 20, 30}

	for i, v := range nums {
		fmt.Printf("index=%d value=%d\n", i, v)
	}

	for _, v := range nums {
		fmt.Println(v)
	}
}

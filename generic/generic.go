package main

import "fmt"

func first[G any](m []G) G {
	return m[0]
}

func second[T any](m []T) T {
	return m[1]
}

func main() {
	number := []int{1, 2, 3, 4, 5}
	strings := []string{"one", "two", "three", "four", "five"}

	fmt.Println(first(number))
	fmt.Println(first(strings))

	fmt.Println(second(number))
	fmt.Println(second(strings))
}

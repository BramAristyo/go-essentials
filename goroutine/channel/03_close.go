package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	// ch <- 10
	// ch <- 20
	// ch <- 30 // in this part will be deadlock the, buffer full dan tidak ada 2 receiver yang aktif

	// go func() {
	ch <- 10
	ch <- 20
	/* 	ch <- 30 */
	// 	// ketika buffer melebihi batas itu di perbolehkan asal ada receiver yang sedang (menunggu value)
	// }()

	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

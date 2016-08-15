package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a = [3]int{8, 1, 2}

	for i, v := range a {
		fmt.Println("%d %d\n", i, v)
	}

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])

	month := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Apr", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dev"}

	Q2 := month[4:7]
	fmt.Println(Q2)
}

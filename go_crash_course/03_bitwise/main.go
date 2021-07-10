package main

import (
	"fmt"
	// "strconv"
	// "math/rand"
	"strings"
)

const (
	UPPER = 1
	LOWER = 2
	CAP = 4 
	REV = 8
)

func main() {
	// fmt.Println(procstr("HELLO PEOPLE!", LOWER|REV|CAP))
	// var x [5]float64
	// x[0] = 98
	// x[1] = 101
	// x[2] = 80
	// x[3] = 76
	// x[4] = 54
	var total float64 = 0
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }
	
	x := [5]float64{ 
		98, 
		101, 
		80, 
		76, 
		54, 
	}
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func procstr(str string, conf byte) string {
	// reverse string
	rev := func(s string) string {
		fmt.Println(s)
		runes := []rune(s)
		fmt.Println(runes)
		n := len(runes)
		for i:=0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}
		return string(runes)
	}

	// query config bits
	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}
	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}
	if (conf & CAP) != 0 {
		str = strings.Title(str)
	}
	if (conf & REV) != 0 {
		str = rev(str)
	}
	return str
}
	// var x int64 = 0xAD
	// var y int64 = 0xF0
	// var z int64 = 160
	// fmt.Println(x)
	// fmt.Println(strconv.FormatInt(0,2))
	// fmt.Println(strconv.FormatInt(x, 2))
	// fmt.Println(strconv.FormatInt(y, 2))

	// fmt.Println(strconv.FormatInt((x & y), 2))
	// fmt.Println(z ^ y)
	// fmt.Println(strconv.FormatInt((z^y), 2))
	// fmt.Println( x & 1 ) 

	// for x:= 0; x < 100; x++ {
	// 	num := rand.Int()
	// 	if num&1 == 1 {
	// 		fmt.Printf("%d is odd\n", num)
	// 	} else {
	// 		fmt.Printf("%d is even\n", num)
	// 	}
	// }

	// var a uint8 = 0
	// a |= 196
	// a |= 3
	// fmt.Printf("%b", a)

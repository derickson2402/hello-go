package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("===== Part 1 =====")
	part1()
	fmt.Println("===== Part 2 =====")
	part2()
	fmt.Println("===== Part 3 =====")
	part3()
	fmt.Println("===== Part 4 =====")
	part4()
	fmt.Println("===== Part 5 =====")
	part5()
	fmt.Println("===== Part 6 =====")
	part6()
	fmt.Println("===== Part 7 =====")
	part7()
}

// Part 1: hello world
func part1() {
	fmt.Println("hi golang!! excited to meet u!")
}

// Part 2: values
func part2() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println("7.0/3 =", 7.0/3)
	fmt.Println("7/3.0 =", 7/3.0)
	fmt.Println("7/3 =", 7/3)
	fmt.Println(true && false)
	fmt.Println("true && false || !false =", true && false || !false)
	fmt.Println("false || !false && true =", false || !false && true)
	fmt.Println("true && !true || !false && true =", true && !true || !false && true)
}

// Part 3: Variables
func part3() {
	var a = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	f := "apple"
	fmt.Println(f)
}

// Part 4: Constants
func part4() {
	const s string = "constant"
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

// Part 5: Looping
func part5() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// Part 6: If/Else
func part6() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Part 7: Switch
func part7() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

package main

import (
	"errors"
	"fmt"
)

const ADULT = 18

func isAdult(age int) bool {
	return age >= ADULT
}

func checkForAdults(ages []int) string {
	for _, age := range ages {
		if isAdult(age) {
			return "Yes"
		}
	}
	return "No"
}

func findUnderageNames(persons map[string]int) (adults []string, err error) {
	for name, age := range persons {
		if !isAdult(age) {
			adults = append(adults, name)
		}
	}
	if len(adults) > 0 {
		err = errors.New("Underages found!")
	}
	return
}

func main() {
	alex := 18
	maxim := 10
	maria := 7

	//Task 1

	fmt.Println("Alex is adult:", isAdult(alex))
	fmt.Println("Maxim is adult:", isAdult(maxim))
	fmt.Println("Maria is adult:", isAdult(maria))

	//Task 2

	fmt.Println(checkForAdults([]int{alex, maria, maxim}))
	fmt.Println(checkForAdults([]int{maxim}))
	fmt.Println(checkForAdults([]int{}))

	//Task 3
	persons := map[string]int{
		"alex":   14,
		"maxim":  18,
		"maria":  10,
		"georgy": 79,
		"mike":   0,
	}

	if adults, err := findUnderageNames(persons); err != nil {
		fmt.Println(err)
		fmt.Println(adults)
	}
	if adults, err := findUnderageNames(map[string]int{}); err != nil {
		fmt.Println(err)
		fmt.Println(adults)
	}
}

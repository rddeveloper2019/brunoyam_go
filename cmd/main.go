package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func IsEven(n int) bool {
	return n%2 == 0 && n != 0
}

func task1() {
	fmt.Println("Задание 1. Чётное или нечётное число?")
	var n int
	fmt.Scan(&n)
	if IsEven(n) {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}
}

func task2() {
	fmt.Println("Задание 2. Найдите максимальное и минимальное числа, среднее арифметическое массива")
	fmt.Println("Введите числа через запятую") //Пытался сделать через пробел, разделение строки не массив не работает корректно
	line := ""
	fmt.Scan(&line)

	words := strings.Split(line, ",") //Пытался сделать через пробел, разделение строки не массив не работает корректно
	var numbers []int
	total := 0
	for _, word := range words {
		number, err := strconv.Atoi(string(word))
		if err == nil {
			numbers = append(numbers, number)
			total += number
		}
	}

	sort.Slice(numbers, func(a, b int) bool { return a < b })
	fmt.Println("Всего чисел:", len(numbers))
	fmt.Println("Сумма:", total)
	fmt.Println("Среднее:", total/len(numbers))
	fmt.Println("Минимальное число:", numbers[0])
	fmt.Println("Максимальное число:", numbers[len(numbers)-1])

}

type Point struct {
	x, y float64
}

func rectanglePerimeter(points []Point) float64 {
	p1 := points[0]
	p2 := points[1]
	length := math.Abs(p2.x - p1.x)
	width := math.Abs(p2.y - p1.y)
	return 2 * (length + width)
}

func task3() {
	//Периметр прямоуголника по двум точкам (математику не знаю, формулу любезно предоставил Chatgpt)
	fmt.Println("Задание 3: по 2 координатам на плоскости, посчитать периметр прямоуголника")

	var points []Point
	var a, b float64

	fmt.Println("Ведите координаты")

	for i := 0; i < 2; i++ {
		fmt.Printf("х:")
		fmt.Scan(&a)
		fmt.Printf("y:")
		fmt.Scan(&b)
		points = append(points, Point{a, b})
	}

	perimeter := rectanglePerimeter(points)
	fmt.Println("Периметр прямоугольника:", perimeter)
}
func main() {
	fmt.Println("_________\n")
	task1()
	fmt.Println("_________\n")
	task2()
	fmt.Println("_________\n")
	task3()
}

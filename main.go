package main

import (
	"fmt"
	"math"
)

func main() {

	// Задача уровня Easy

	var a int

	fmt.Println("Введите число:")
	fmt.Scan(&a)

	remainder := a % 2
	rez := remainder == 0

	fmt.Println("Число четное:", rez)

	// Задача уровня Normal

	mas := [5]int{3, 4, 4, 7, 1}

	fmt.Println("Массив:", mas)

	max := mas[0]
	min := mas[0]
	sum := 0

	for i := 0; i < len(mas); i++ {

		if mas[i] > max {

			max = mas[i]

		}

		if mas[i] < min {

			min = mas[i]

		}

		sum = sum + mas[i]

	}

	var average float64
	average = float64(sum) / float64(len(mas))

	fmt.Println("Максимальное число:", max)
	fmt.Println("Минимальное число:", min)
	fmt.Println("Среднее арифметическое:", average)

	//Задача уровня Hard
	//Формула площади треугольника по координатам вершин:
	//S = 0.5*|(x2 - x1)(y3-y1) - (x3-x1)(y2-y1)|

	masX := [3]float64{3, 4, 5}
	masY := [3]float64{-5, 3.5, 7}

	fmt.Println("Координаты точек на оси x:", masX)
	fmt.Println("Координаты точек на оси y:", masY)

	var S float64

	S = 0.5 * math.Abs((masX[1]-masX[0])*(masY[2]-masY[0])-(masX[2]-masX[0])*(masY[1]-masY[0]))

	fmt.Println("Площадь треугольника:", S)

	//Конец

}

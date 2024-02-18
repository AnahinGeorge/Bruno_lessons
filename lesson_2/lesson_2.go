package lesson_2

import (
	"errors"
	"fmt"
)

func TestEasy() {

	// Задача уровня Easy

	var age int

	fmt.Println("Введите возраст")
	fmt.Scan(&age)

	res := SetAge(age)

	if res == true {

		fmt.Println("Человек совершеннолетний")

	} else {

		fmt.Println("Человек несовершеннолетний")

	}

}

func SetAge(age int) (res bool) {

	if age < 18 {

		res = false

	} else {

		res = true

	}

	return res

}

func TestNormal() {

	//Задача уровня Normal

	var ageSlice = make([]int, 5, 10)

	fmt.Println("Введите возраста для 5 человек")

	for i := 0; i < len(ageSlice); i++ {

		fmt.Scan(&ageSlice[i])

	}

	res := HaveAdult(ageSlice)

	if res == true {
		fmt.Println("В наборе есть совершеннолетний")
	} else {

		fmt.Println("В наборе нет совершеннолетних")

	}
}

func HaveAdult(ageSlice []int) (res bool) {

	res = false

	for _, age := range ageSlice {

		if age >= 18 {

			res = true
			break

		}

	}

	return res

}

func TestHard() {

	// Задача уровня Hard

	ageMap := make(map[string]int)
	var quantity int
	var name string
	var age int

	fmt.Println("Введите количество человек:")
	fmt.Scan(&quantity)

	for i := 0; i < quantity; i++ {

		fmt.Println("Введите имя")
		fmt.Scan(&name)
		fmt.Println("Введите возраст")
		fmt.Scan(&age)
		ageMap[name] = age

	}

	ReturnNames(ageMap)

}

func ReturnNames(ageMap map[string]int) {

	var nameSlice []string

	for kay, value := range ageMap {

		if value < 18 {

			nameSlice = append(nameSlice, kay)

		}

	}

	if len(nameSlice) > 0 {

		err := errors.New("ЕСТЬ НЕСОВЕРШЕННОЛЕТНИЕ")
		fmt.Println(nameSlice)
		fmt.Println(err)

	}

}

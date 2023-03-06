package main

import (
	"fmt"
)

func main() {
	exercise26()
}

func exercise17() {
	// Change value at position 1 by 'cucumber'
	card := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	card[1] = "cucumber"
	fmt.Print(card)
}

func exercise18() {
	// Change value at position 1 by 'cucumber'
	card := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	for i, v := range card {
		fmt.Println(i, " -> ", v)
	}
}

func exercise20() {
	// Add 'olive-oil'
	card := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	card = append(card, "olive-oli")
	fmt.Print(card)
}

func exercise21() {
	// Join both arrays
	treballadors := []string{"Josep", "Cristina", "Helena", "Robert"}
	novesIncorporacions := []string{"Ivan", "Estel", "Aitana"}
	treballadors = append(treballadors, novesIncorporacions...)
	fmt.Print(treballadors)
}

func exercise22() {
	myArray := make([]int, 9, 10)
	fmt.Printf("The cap is: %d and the len is: %d\n", cap(myArray), len(myArray))
	myArray = append(myArray, 1)
	fmt.Printf("The cap is: %d and the len is: %d\n", cap(myArray), len(myArray))
	myArray = append(myArray, 2)
	fmt.Printf("The cap is: %d and the len is: %d\n", cap(myArray), len(myArray))
}

func exercise23() {
	// Create a map of three students and their qualifications
	myMap := map[string]int{
		"Student1": 5,
		"Student2": 6,
		"Student3": 4,
	}
	fmt.Println(myMap)
}

func mapexercise() {
	myMap := map[string]int{
		"Student1": 5,
		"Student2": 6,
		"Student3": 4,
	}

	if v, ok := myMap["Student4"]; ok {
		fmt.Println("S’ha borrat l’element amb el valor", v, ok)
		delete(myMap, "Student4")
	} else {
		fmt.Println("No existeix l’element amb el valor", v, ok)
	}
}

func exercise26() {
	//A continuació realitza 2 structs anonims sobre articles de supermercats a on tens que definir les dades de nom, unitats i pvp.
	article := struct {
		name  string
		units int
		pvp   int
	}{
		name:  "name1",
		units: 2,
		pvp:   4,
	}

	fmt.Println(article)

}

package main

import (
	"fmt"
	"reflect"
	//"math/rand"
)

func main() {

	var stringType string
	var intType int
	var intArray [5]int
	var floatType float64
	var booleanType bool
	arr := []interface{}{stringType, intType, floatType, booleanType}

	types := fmt.Sprint("The types of ", stringType, intType, floatType, booleanType, arr, " are the next..")
	fmt.Println(types)
	fmt.Printf("%T\n", stringType)
	fmt.Printf("%T\n", intType)
	fmt.Printf("%T\n", intArray)
	fmt.Printf("%T\n", floatType)
	fmt.Printf("%T\n", arr)

	fmt.Println("\nTypeOf by: reflect.TypeOf(var)")
	fmt.Println(reflect.TypeOf(stringType))
	fmt.Println(reflect.TypeOf(intType))
	fmt.Println(reflect.TypeOf(floatType))
	fmt.Println(reflect.TypeOf(booleanType))
	fmt.Println(reflect.TypeOf(arr))

	fmt.Println("\nTypeOf by: reflect.TypeOf(var).Kind()")
	fmt.Println(reflect.TypeOf(stringType).Kind())
	fmt.Println(reflect.TypeOf(intType).Kind())
	fmt.Println(reflect.TypeOf(floatType).Kind())
	fmt.Println(reflect.TypeOf(booleanType).Kind())
	fmt.Println(reflect.TypeOf(arr).Kind())

	/*
		arr := []interface{}{stringType, intType, floatType, booleanType}
		fmt.Println(arr)
		showVarTypes(arr)
	*/
}

func showVarTypes(arr interface{}) {
	//for pos := 0; pos <= ; pos++ {
	//}

}

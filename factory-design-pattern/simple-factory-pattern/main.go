package main

import "fmt"

func main() {
	//create a new service logistics with 20 tons
	sL := GetServiceLogicstics(20)
	fmt.Println("Level demo simple factory design", sL)

	//create a new service logistics with default values
	sL2 := GetServiceLogicstics(10)
	fmt.Println("Level demo simple factory design", sL2)
}

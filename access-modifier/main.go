package main

import (
	"access-modifier/library"
	"fmt"
)

func main() {
	fmt.Println("HourInADay:", library.HourInADay)
	// fmt.Println("secondsInMinute:", library.secondsInMinute)
	person1 := library.Person{"Beni", 24}
	fmt.Println("person1: ", person1)
	fmt.Println("Name:", library.Name)
	fmt.Println("25 Hari ke menit:", library.DaysToMinute(25))
}

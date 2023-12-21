package main

import "fmt"

func main() {
	user := map[string]string{
		"username": "benedictJ",
		"email":    "benedictsujullian@gmail.com",
	}

	fmt.Println(user)

	fmt.Println("=======================================")

	scores := make(map[string]int)
	fmt.Println(scores)

	scores["java"] = 85
	scores["react"] = 87
	scores["kotlin"] = 90

	fmt.Println("Scores: ", scores)
	fmt.Println("Nilai Java: ", scores["java"])
	fmt.Println("Nilai Kotlin: ", scores["kotlin"])
	fmt.Println("Nilai Golang: ", scores["golang"]) // will get default value if the key has no value, 0 for int
	fmt.Println()

	fmt.Println("==========reassign map value===========")

	scores["java"] = 90

	fmt.Println("Scores: ", scores)
	fmt.Println()

	fmt.Println("==========delete value in map===========") // delete(map, key)

	delete(scores, "kotlin")
	fmt.Println("Scores: ", scores)

	fmt.Println("===============looping map==============")

	names := map[int]string{
		1: "john",
		2: "jane",
		3: "lily",
		4: "ruby",
	}

	for key, value := range names {
		fmt.Println("Key,   :", key)
		fmt.Println("value, :", value)
	}
}

package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}
func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}

	person, ok := personDB["12345"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 12345")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}

	//myMap := make(map[string]PersonInfo)
	//myMap1 := make(map[string]PersonInfo, 100)

	myMap2 := map[string]PersonInfo{
		"1234": PersonInfo{"1", "Jack", "Room 101,..."},
	}
	fmt.Println(myMap2)
	fmt.Println(example(1))

	/*
		for j := 0; j < 5; j++ {
			for i := 0; i < 10; i++ {
				if i > 5 {
					break JLoop
				}
				fmt.Println(i)
			}
		}
		JLoop:
		fmt.Println("finish")
	*/

}

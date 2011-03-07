package main

import (
	. "./maps"
	"fmt"
)


//---
type Person struct {
	Name    string
	Surname string
	Age     int
}

func (p Person) String() string {
	s := p.Name + "#" + p.Surname
	return s
}
//---


func main() {

    printValues := func (v Pair) {
	    fmt.Printf("Person, Job = %v, %s\n", v.Key, v.Value)
    }

	p1 := Person{"John", "Rambo", 60}
	p2 := Person{"John", "Doe", 30}
	p3 := Person{"Jane", "Doe", 30}

	//Create map
	M := NewSMap()

	//Insert into map
	M.Insert(p1, "Manager")
	M.Insert(p2, "Employee")
	M.Insert(p3, "Secretary")

	//Iterate Map
	M.Do(printValues)

	println("---")

	//Query map
	p := Person{"John", "Rambo", 60}
	result := M.Get(p)
	if result.Ok {
		fmt.Printf("Job = %s\n", result.Value)
	}

	println("---")

	//Delete from map
	M.Delete(p1)

	//check
	M.Do(printValues)

}

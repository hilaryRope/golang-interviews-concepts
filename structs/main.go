package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type SecretAgent struct {
	Person
	ltk bool
}

func main() {
	p1 := Person{firstName: "John", lastName: "Snow", age: 24}
	p2 := Person{firstName: "James", lastName: "Bond", age: 32}

	sa := SecretAgent{p2, true}

	fmt.Println("Person 1 ", p1.firstName, p1.lastName, p1.age)
	fmt.Println("Person 2 ", p2.firstName, p2.lastName, p2.age)

	fmt.Println("SecretAgent ", sa.firstName, sa.lastName, sa.age, sa.ltk)

	fmt.Printf("%T \t %#v", p1, p2)
	fmt.Printf("%T \t", sa)

	// Anonymous structs
	ap := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Peter", lastName: "Pan", age: 12,
	}

	fmt.Printf("%T\n", ap)
	fmt.Printf("First Name: %s, Last Name: %s, Age: %d\n", ap.firstName, ap.lastName, ap.age)

}

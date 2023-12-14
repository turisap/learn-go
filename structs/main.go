package main

import (
	"fmt"
	"strings"
)

type book struct {
	title  string
	author string
	year   int
}

// if fields are of the same type
type book1 struct {
	title, author string
	year          int
}

func main() {
	//basics()
	embedded()
}

func embedded() {
	fmt.Println("\n" + strings.Repeat("@", 20))

	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John K.",
		salary: 60000,
		contactInfo: Contact{
			email:   "jk@mail.com",
			address: "Batumi",
			phone:   33224411,
		},
	}

	fmt.Printf("%#v\n", john)
	fmt.Println(john.contactInfo.email)
	john.contactInfo.email = "dksfj"
	fmt.Println(john.contactInfo.email)
}

func basics() {

	myBook := book{
		title:  "the middle passage",
		author: "Hollis",
		year:   1345,
	}

	fmt.Println(myBook)

	anotherBook := book{
		title: "kis",
	}

	// {title:"kis", author:"", year:0}
	// zero-initialized
	fmt.Printf("%#v", anotherBook)

	type user struct {
		name, email, password string
		balance               float64
	}

	u := user{name: "kirill"}

	// accessing non-existing field
	// will raise an error
	//fmt.Println(u.trololo)

	fmt.Println(u.name)

	u.email = "k@mgmail.com"

	fmt.Printf("%#v", u)
	fmt.Println("\n" + strings.Repeat("@", 20))

	u1 := user{"kirill", "k@gmail.com", "sdkfh", 0.4}
	u2 := user{"kirill", "k@gmail.com", "sdkfh", 0.4}

	// can we compare two structs?
	// prints out true
	fmt.Println(u1 == u2)

	// cloning structs
	user1 := user{"anna", "a@gmail", "rijsldjf", 0.3}
	// this creates a deep copy
	user2 := user1

	user2.name = "TROLOLOSH"

	// prints out {name:"anna", ...}
	fmt.Printf("%#v\n", user1)
	// prints out {name:"TROLOLOSH", ...}
	fmt.Printf("%#v\n", user2)

	fmt.Println("\n" + strings.Repeat("@", 20))

	// declaration of an anonymous struct
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "diana",
		lastName:  "Kis",
		age:       33,
	}

	fmt.Printf("%#v", diana)

	// anonymous struct
	type abook struct {
		string
		float64
	}

	ab := abook{"The midpassage", 33.3}

	// prints {string:"The midpassage", float64:33.3}
	fmt.Printf("%#v", ab)

}

package stripper

import (
	"fmt"
	"log"

	"github.com/hawry/stripper"
)

func ExampleMarshal() {
	type User struct {
		Username string `json:"username"`
		Password string `json:"password" clean:"true"`
	}
	user := User{
		"hawry",
		"abadpassword",
	}
	json, err := stripper.Marshal(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", json)
	// Output: {"username":"hawry","password":""}
}

func ExampleMarshalIndent() {
	type User struct {
		Username string `json:"username"`
		Password string `json:"password" clean:"true"`
	}
	user := User{
		"hawry",
		"abadpassword",
	}
	json, err := stripper.MarshalIndent(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", json)
	// Output: {
	//   "username": "hawry",
	//   "password": ""
	// }
}

func ExampleClean() {
	type User struct {
		Username string `json:"username"`
		Password string `json:"password" clean:"true"`
	}
	user := User{
		"hawry",
		"abadpassword",
	}
	mod, err := stripper.Clean(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", mod)
	// Output: &{Username:hawry Password:}
}

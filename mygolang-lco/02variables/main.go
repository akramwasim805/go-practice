package main

import "fmt"

const LOGIN_TOKEN = "edgdf434343ddcdscvsd" // this is a public one

func main() {

	fmt.Println("i am using some Variables here in this module")

	//implicit
	var someInt = 5
	var someString = "hi i am wasim"
	fmt.Println(someString+", hey hwo are you we are testing some thing", someInt, "hi i can give any option here")
	fmt.Printf("This is formattor and it will print my string here %v and my int here %v \n", someString, someInt)

	// no var style

	numberOfUser := 300
	nameOfTeacher := "wasim"

	fmt.Printf("This is formattor and it will print my string here %v and my int here %v \n", nameOfTeacher, numberOfUser)

}

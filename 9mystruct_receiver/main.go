package main

import "fmt"

type Username struct {
	email string
}

func (usr *Username) Login() string {
	return fmt.Sprintf("Login %v", usr.email)
}

type Employee struct {
	firstname string
	salary    int
}

func (emp *Employee) Login() string {
	return fmt.Sprintf("Login %v", emp.firstname)
}

type Activity interface {
	Login() string
}

func main() {

	var login Activity = &Employee{}
	login = &Username{email: "erhanburhanudin@gmail.com"}
	fmt.Println(login.Login())

}

package contact

import "fmt"

// Get is a function that returns the contact details
func Read() any {
	fmt.Println("get contact")
	return "get contact"
}
func Create() any {
	fmt.Println("create contact")
	return "create contact"
}
func Delete() any {
	fmt.Println("delet contact")
	return "delete contact"
}
func Update() any {
	fmt.Println("update contact")
	return "update contact"
}

// Busca conta
func Search() any {
	fmt.Println("saerch contact")
	return "search contact"
}

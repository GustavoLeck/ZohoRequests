package account

import "fmt"

// Get is a function that returns the account details
func Read() any {
	fmt.Println("get account")
	return "get account"
}
func Create() any {
	fmt.Println("create account")
	return "create account"
}
func Delete() any {
	fmt.Println("delet account")
	return "delete account"
}
func Update() any {
	fmt.Println("update account")
	return "update account"
}

// Busca conta
func Search() any {
	fmt.Println("saerch account")
	return "search account"
}

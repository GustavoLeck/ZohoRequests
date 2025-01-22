package account
import "fmt"
//Get is a function that returns the account details
func Read() any {
	fmt.Println("get")
	return "get"
}
func Create() any {
	fmt.Println("create")
	return "create"
}
func Delete() any {
	fmt.Println("delete")
	return "delete"
}
func Update() any {
	fmt.Println("update")
	return "update"
}
//Busca conta
func Search() any {
	fmt.Println("saerch")
	return "search"
}

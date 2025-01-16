package account
import "fmt"
//Get is a function that returns the account details
func read() any {
	fmt.Println("get")
	return "get"
}

func create() any {
	fmt.Println("create")
	return "create"
}

func delete() any {
	fmt.Println("delete")
	return "delete"
}

func update() any {
	fmt.Println("update")
	return "update"
}

//Busca conta
func search() any {
	fmt.Println("saerch")
	return "search"
}

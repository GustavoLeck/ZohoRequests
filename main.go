package main

import (
	"fmt"
	"net/http"

	"github.com/GustavoLeck/ZohoRequests/crm/CrmToken"
)

func main() {
	CrmToken.SetParameters("1000.QN3NAUYJWDCNP8U8SQEQCNPV34A46W", "a6ae958f0793193ec7fc7d95423ae247bbd0fbbd24", "1000.86dacc2e6446f9828360b1a066ceb0fd.61ad74ae521b20f17aeab0bf193ea542")
	HandleRequest()
}

func HandleRequest() {
	fmt.Println("Server is running on port 8000")
	// http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}

// func Home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Print(w, "Home Page")
// }

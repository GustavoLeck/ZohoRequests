package main

import (
	"fmt"

	"github.com/GustavoLeck/ZohoRequests/CRM/CrmToken"
)

func main() {
	fmt.Println("Teste")
	//1000.86dacc2e6446f9828360b1a066ceb0fd.61ad74ae521b20f17aeab0bf193ea542
	CrmToken.SetParameters("1000.QN3NAUYJWDCNP8U8SQEQCNPV34A46W", "a6ae958f0793193ec7fc7d95423ae247bbd0fbbd24", "1000.86dacc2e6446f9828360b1a066ceb0fd.61ad74ae521b20f17aeab0bf193ea542")
	fmt.Println(CrmToken.AcessToken)
	CrmToken.UpdateAcessToken()
	fmt.Println(CrmToken.AcessToken)
}

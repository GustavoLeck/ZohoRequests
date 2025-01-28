package ZohoCrmAccount

import (
	"bytes"
	"fmt"
	"net/http"
)

type version float64

func selectVersion(version version) string {
	switch version {
	case 2:
		return "v2"
	case 2.1:
		return "v2.1"
	case 3:
		return "v3"
	case 4:
		return "v4"
	case 5:
		return "v5"
	case 6:
		return "v6"
	case 7:
		return "v7"
	default:
		panic("version api not found")
	}
}

/*
This function create an account from the Zoho CRM database
*/
func Create(version version, idAccount string, value map) (*http.Response, error) {
	var versionsSelected string = selectVersion(version)
	var maxRetries int = 5
	var retries int = 0
	for {
		response, err := http.Post("https://www.zohoapis.com/crm/"+versionsSelected+"/Accounts/"+idAccount, "application/json", value)
		if err != nil {
			return response, nil
		}
		if maxRetries == retries {
			return nil, err
		}
		retries++
	}
}

/*
This function get an account from the Zoho CRM database
*/
func Read(version version, idAccount string) (*http.Response, error) {
	var versionsSelected string = selectVersion(version)
	var maxRetries int = 5
	var retries int = 0
	for {
		response, err := http.Get("https://www.zohoapis.com/crm/" + versionsSelected + "/Accounts/" + idAccount)
		if err != nil {
			return response, nil
		}
		if maxRetries == retries {
			return nil, err
		}
		retries++
	}
}

/*
This function udpate an account from the Zoho CRM database
*/
func Update(version float64, idAccount string, value map) (*http.Response, error) {
	var versionsSelected string = selectVersion(version)
	var maxRetries int = 5
	var retries int = 0
	for {
		response, err := http.MethodPut("https://www.zohoapis.com/crm/" + versionsSelected + "/Accounts/" + idAccount, value)
		if err != nil {
			return response, nil
		}
		if maxRetries == retries {
			return nil, err
		}
		retries++
	}
}


/*
This function deletes an account from the Zoho CRM database
*/
func Delete(version version, idAccount string) (*http.Response, error) {
	var versionsSelected string = selectVersion(version)
	var maxRetries int = 5
	var retries int = 0
	for {
		response, err := http.MethodDelete("https://www.zohoapis.com/crm/" + versionsSelected + "/Accounts/" + idAccount)
		if err != nil {
			return response, nil
		}
		if maxRetries == retries {
			return nil, err
		}
		retries++
	}
}

/*
Search is a function that returns the account details, using a criteria to filter registers in the Zoho CRM database
https://www.zoho.com/deluge/help/crm/search-records.html
*/
func Search(version version, criteria string)  (*http.Response, error) {
	var versionsSelected string = selectVersion(version)
	var maxRetries int = 5
	var retries int = 0
	for {
		response, err := http.MethodDelete("https://www.zohoapis.com/crm/" + versionsSelected + "/Accounts/", criteria)
		if err != nil {
			return response, nil
		}
		if maxRetries == retries {
			return nil, err
		}
		retries++
	}
}

package ZohoCrmAccount

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func selectVersion(version float64) string {
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
		panic("version api selected not found")
	}
}

func convertReader(value map[string]interface{}) (*bytes.Reader, error) {
	valueJson, errJson := json.Marshal(value)
	if errJson != nil {
		return nil, errJson
	}
	return bytes.NewReader(valueJson), nil
}

/*
This function create an account from the Zoho CRM database
*/
func Create(version float64, idAccount string, value map[string]interface{}) (*http.Response, error) {
	var versionsSelected string = selectVersion(version)
	valueConverted, valueErr := convertReader(value)
	if valueErr != nil {
		return nil, valueErr
	}
	var maxRetries int = 5
	var retries int = 0
	for {
		response, err := http.Post("https://www.zohoapis.com/crm/"+versionsSelected+"/Accounts/"+idAccount, "application/json", valueConverted)
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
func Read(version float64, idAccount string) (*http.Response, error) {
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
func Update(version float64, idAccount string, value map[string]interface{}) (*http.Response, error) {
	var versionsSelected string = selectVersion(version)
	valueConverted, valueErr := convertReader(value)
	if valueErr != nil {
		return nil, valueErr
	}
	var maxRetries int = 5
	var retries int = 0
	for {
		rep, err := http.NewRequest(http.MethodPut, "https://www.zohoapis.com/crm/"+versionsSelected+"/Accounts/"+idAccount, valueConverted)
		if err != nil {
			return nil, err
		}
		client := &http.Client{}
		response, err := client.Do(req)
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
func Delete(version float64, idAccount string) (*http.Request, error) {
	var versionsSelected string = selectVersion(version)
	var maxRetries int = 5
	var retries int = 0
	for {
		response, err := http.NewRequest(http.MethodDelete, "https://www.zohoapis.com/crm/"+versionsSelected+"/Accounts/"+idAccount, nil)
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
func Search(version float64, criteria string) (*http.Response, error) {
	var versionsSelected string = selectVersion(version)
	var maxRetries int = 5
	var retries int = 0
	for {
		rep, err := http.Get("https://www.zohoapis.com/crm/" + versionsSelected + "/Accounts?criteria=" + criteria)
		if err != nil {
			return rep, nil
		}
		if maxRetries == retries {
			return nil, err
		}
		retries++
	}
}

package CrmToken

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var cliId string
var cliSecret string
var rToken string
var AcessToken string

func SetParameters(clientId string, clientSecret string, refreshToken string) {
	cliId = clientId
	cliSecret = clientSecret
	rToken = refreshToken
	UpdateAcessToken()
}

func UpdateAcessToken() {
	res, err := http.Post("https://accounts.zoho.com/oauth/v2/token?grant_type=refresh_token&refresh_token="+rToken+"&client_id="+cliId+"&client_secret="+cliSecret, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Erro ao gerar token: " + err.Error())
		return
	}
	defer res.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		fmt.Println("Erro ao decodificar o JSON: " + err.Error())
		return
	}
	if accessToken, ok := result["access_token"].(string); ok {
		AcessToken = accessToken
	} else {
		fmt.Println("Erro: access_token não encontrado na resposta")
	}
}

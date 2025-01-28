package ZohoToken

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/robfig/cron"
)

var cliId string
var cliSecret string
var rToken string
var acToken string

func SetParameters(clientId string, clientSecret string, refreshToken string) {
	cliId = clientId
	cliSecret = clientSecret
	rToken = refreshToken
	UpdateAcessToken()
	c := cron.New()
	c.AddFunc("40 * * * *", func() {
		fmt.Println("Atualizando o Acess Token", time.Now().Format("2006-01-02 15:04:05"))
		UpdateAcessToken()
	})
	c.Start()
}

func UpdateAcessToken() {
	res, err := http.Post("https://accounts.zoho.com/oauth/v2/token?grant_type=refresh_token&refresh_token="+rToken+"&client_id="+cliId+"&client_secret="+cliSecret, "application/x-www-form-urlencoded", nil)
	if err != nil {
		panic("Erro ao tentar obter o access_token: " + err.Error())
	}
	defer res.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		panic("Erro ao decodificar o JSON: " + err.Error())
	}
	if accessToken, ok := result["access_token"].(string); ok {
		acToken = accessToken
	}
}

func GetToken() (string, error) {
	if acToken == "" {
		panic("AcessToken não foi gerado")
	}
	return acToken, nil
}

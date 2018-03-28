package service

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Token defines the token content
type Token struct {
	IDToken      string `json:"ID_Token"`
	RefreshToken string `json:"Refresh_Token"`
	Expiration   string `json:"Expiry"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func genericLookup(url, username, password string) interface{} {

	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))

	if err != nil {
	}
	resp, err3 := client.Do(req)

	if err3 != nil {
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			return err2
		}
		bodyString := string(bodyBytes)
		var data interface{}
		json.Unmarshal([]byte(bodyString), &data)
		return data
	}
	return nil

}

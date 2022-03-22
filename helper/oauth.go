package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type AccessTokenJSON struct {
	AccessToken string `json:"accessToken"`
}

func GetUniAuthAccessToken(code string) (string, error) {
	uniAuthBaseUrl := viper.GetString("uniAuthBaseUrl")
	clientSecret := viper.GetString("uniAuthSecret")
	uniAuthClientId := viper.GetString("uniAuthClientId")
	data := make(map[string]interface{})
	data["clientId"] = uniAuthClientId
	data["clientSecret"] = clientSecret
	data["code"] = code
	bytesData, _ := json.Marshal(data)
	resp, err := http.Post(uniAuthBaseUrl+"/api/oauth/access-token", "application/json", bytes.NewReader(bytesData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	accessTokenJSON := &AccessTokenJSON{}
	json.Unmarshal(body, accessTokenJSON)
	if accessTokenJSON.AccessToken == "" {
		return "", fmt.Errorf("授权登录失败，请重试")
	}
	return accessTokenJSON.AccessToken, nil
}
